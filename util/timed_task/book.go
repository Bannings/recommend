package timed_task

import (
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"strings"
	"time"
)

func getBookTag() []model.CartoonTag {
	var tagData []model.CartoonTag
	err := db.RecallDB().Model(&model.CartoonTag{}).Find(&tagData).Error
	if err != nil {
		log.Errorf("get cartoon tag fail:%v", err)
	}
	return tagData
}

func getTagMap(cartoonTags []model.CartoonTag) map[int]model.TagInfo {
	tagMap := make(map[int]model.TagInfo)
	for _, cartoonTag := range cartoonTags {
		category := ""
		if cartoonTag.Category != "NULL" {
			category = cartoonTag.Category
		}
		tagInfo := model.TagInfo{Tag: cartoonTag.Tag, Category: category}
		tagMap[cartoonTag.Cid] = tagInfo
	}
	return tagMap
}

func getComicTips() {
	var ComicFrees string

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc).Format("2006-01-02")

	err := db.ComicDataDB().Table("cartoon_limit_pay").Select("ids").Where("date_flag=?", now).Limit(1).Row().Scan(&ComicFrees)
	if err != nil {
		log.Errorf("get cartoon tag fail:%v", err)
		return
	}
	comics := strings.Split(string(ComicFrees), ",")
	descriptionMap := make(map[string]int, len(comics))
	for i := range comics {
		descriptionMap[comics[i]] = 1
	}
	model.RecommendMaps.ComicFreeMap = descriptionMap
}

func LoadNewDescription() {
	var cartoonDescriptions []model.CartoonDescription

	err := db.ComicDataDB().Model(&model.CartoonDescription{}).Limit(200).Find(&cartoonDescriptions).Error
	if err != nil {
		log.Errorf("get cartoon tag fail:%v", err)
		return
	}
	descriptionMap := make(map[string]string, len(cartoonDescriptions))
	for i := range cartoonDescriptions {
		descriptionMap[cartoonDescriptions[i].CartoonId] = cartoonDescriptions[i].Title
	}
	model.RecommendMaps.DescriptionMap = descriptionMap
}

func UpdateTagMap() {
	getComicTips()
	tags := getBookTag()
	tagMap := getTagMap(tags)
	model.RecommendMaps.CartoonTagMap = tagMap
}
