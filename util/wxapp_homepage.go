package util

import (
	"github.com/xndm-recommend/go-utils/type_"
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"recommend/util/recommend_api"
	"recommend/util/timed_task"
	"strconv"
)

func HomepageTopWxappApplyV2(userInfo model.UserInfo) []model.SectionDataConfig {

	var recommendAll []string
	var topColumns []string
	topColumns = golbal.WxappTolColumns

	recommendIds := make(chan string, len(topColumns)*6)
	topCols := make([]model.SectionDataConfig, len(topColumns), len(topColumns)+golbal.ComicListNum+1)

	for i, sectionId := range topColumns {
		if golbal.INDEX_BANNER == sectionId {
			bannerSection := timed_task.GetBannerSection(userInfo, userInfo.DeviceType, sectionId)
			topCols[i] = *bannerSection
		} else {
			recommendList := recommend_api.GetTopRecommendWxapp(userInfo, sectionId)
			for _, comicId := range recommendList {
				recommendIds <- comicId
			}
			topCols[i] = recommend_api.GetWxappColumnSection(recommendList, sectionId)
		}
	}

	close(recommendIds)
	for comicId := range recommendIds {
		recommendAll = append(recommendAll, comicId)
	}

	// redis 数据更新
	uidStr := strconv.FormatInt(userInfo.Uid, 10)
	key := db.PersonalRead.ItemGetKey(uidStr, userInfo.UdId, type_.IntToStr(userInfo.Gender), golbal.COMIC_SHOW)
	err := db.PersonalRead.ItemZAdd(recommendAll, key)
	if err != nil {
		log.Error(err)
	}
	userInfo.UdId = removeIPUdid(userInfo.UdId)
	return topCols
}

func WxappHomepageMoreApplyV2(userInfo model.UserInfo, sectionId, abTestInd string, pageNum, pageSize int) []model.SectionDataConfig {

	var recommendList []model.SectionDataConfig
	sectionIDInt, _ := strconv.Atoi(sectionId)

	switch {
	case sectionIDInt >= golbal.NEW_USER_COMIC_LIST:
		{
			outComicList := internal_api.GetNewUserComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	case sectionIDInt >= golbal.DYNAMIC_COMIC_LIST:
		{
			outComicList := internal_api.GetNewUserComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	default:
		{
			outComicList := recommend_api.GetTopRecommendMore(userInfo, sectionId, pageNum, pageSize)
			section := recommend_api.GetWxappColumnSection(outComicList, sectionId)
			recommendList = []model.SectionDataConfig{section}
		}
	}

	return recommendList
}
