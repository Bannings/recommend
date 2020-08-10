package timed_task

import (
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"strconv"
	"strings"
)

//老版本banner逻辑
//针对V3接口以下
//新版banner参考运营banner
//internal_api.GetOperationInfo(userInfo, "HomeTOPBanner")
//recommend_api.GetOperationBannerSection(operationinfo, sectionId)

func addBannerData(genderId int) []model.BannerInfo {
	var bannerData []model.BannerInfo
	err := db.BannerDB().Model(&model.BannerInfo{}).Where("status=1 and gender= ? and NOW()>=start_time and NOW()<=end_time", genderId).Find(&bannerData).Error
	if err != nil {
		log.Error(err)
	}
	return bannerData
}

func checkBanner(banner []model.BannerInfo, min, sGender int) []model.BannerInfo {
	banner = removeDuplicate(banner)
	if len(banner) < min {
		return defaultBanner(sGender)
	}
	return banner
}

//针对图片地址去重
func removeDuplicate(banner []model.BannerInfo) []model.BannerInfo {
	var result []model.BannerInfo
	for i := range banner {
		flag := true
		for j := range result {
			if banner[i].ImgUrl == result[j].ImgUrl {
				flag = false
				break
			}
		}
		if flag {
			temp := urlCheck(banner[i].ImgUrl)
			banner[i].ImgUrl = temp
			result = append(result, banner[i])
		}
	}
	return result
}

func ReloadBannerData() {
	minSize := 4
	bannerOut := make([][]model.BannerInfo, 2)
	// 男版banner

	maleBanner := addBannerData(golbal.MALE_ID)
	bannerOut[golbal.MALE_ID] = checkBanner(maleBanner, minSize, golbal.MALE_ID)

	// 女版banner
	femaleBanner := addBannerData(golbal.FEMALE_ID)
	bannerOut[golbal.FEMALE_ID] = checkBanner(femaleBanner, minSize, golbal.FEMALE_ID)

	model.Banner.BannerDatas = bannerOut
}

//默认banner，仅banner个数不足2会使用
func defaultBanner(sGender int) []model.BannerInfo {
	info1 := model.BannerInfo{ComicId: 107182, ComicName: "君心劫", ImgUrl: "cms/chendan/48698d20-4ea8-11e9-ba02-8df2471819ce.png", Platform: 0}
	info2 := model.BannerInfo{ComicId: 106902, ComicName: "我的男神是仓鼠", ImgUrl: "cms/chendan/f8407990-54f0-11e9-a4cd-317e347526c5.png", Platform: 0}
	info3 := model.BannerInfo{ComicId: 105905, ComicName: "明星是血族", ImgUrl: "cms/chendan/6bdfcce0-718c-11e9-8e49-dd74d82de68c.png", Platform: 0}
	info4 := model.BannerInfo{ComicId: 105618, ComicName: "冷酷总裁的夏天", ImgUrl: "cms/chendan/513fec00-67c6-11e9-89ce-bfa4b04495af.png", Platform: 0}
	info5 := model.BannerInfo{ComicId: 107470, ComicName: "我是玉皇大帝", ImgUrl: "cms/chendan/7558f180-aa11-11e9-8fb5-b579ec4d4879.png", Platform: 0}
	info6 := model.BannerInfo{ComicId: 106992, ComicName: "不完美游戏", ImgUrl: "cms/chendan/ada47810-5fe4-11e9-a336-efbe23ec873b.png", Platform: 0}
	info7 := model.BannerInfo{ComicId: 107682, ComicName: "仙武帝尊", ImgUrl: "cms/chendan/66357ba0-1713-11ea-b071-89ea94b7f4b3.jpg", Platform: 0}
	info8 := model.BannerInfo{ComicId: 86080, ComicName: "斗罗大陆3龙王传说", ImgUrl: "cms/chendan/044caa20-662f-11e9-bf34-1396bddeb577.png", Platform: 0}
	banner1 := []model.BannerInfo{info1, info2, info3, info4}
	banner2 := []model.BannerInfo{info5, info6, info7, info8}
	if sGender == golbal.MALE_ID {
		return banner2
	} else if sGender == golbal.FEMALE_ID {
		return banner1
	} else {
		return banner2
	}
}

func GetBannerSection(userInfo model.UserInfo, platform int, sectionId string) *model.SectionDataConfig {

	bannerInfo := make([]model.ComicInfo, 0, 10)

	for _, info := range model.Banner.BannerDatas[userInfo.Gender] {
		if info.Platform == golbal.BANNER_PLATFORM_ALL || platform == info.Platform {
			var comicInfo model.ComicInfo
			comicInfo.ComicID = strconv.Itoa(info.ComicId)
			comicInfo.ComicName = info.ComicName
			comicInfo.ImgUrl = info.ImgUrl
			comicInfo.Url = info.Url
			bannerInfo = append(bannerInfo, comicInfo)
		}
	}
	section := model.BannerSection
	section.ComicInfo = bannerInfo
	section.SectionId = sectionId
	section.SectionName = golbal.ColumnNameMap[sectionId]
	return section
}

func urlCheck(url string) string {
	if strings.Contains(url, "!webp!webp") {
		log.Errorf("wrong format url %s", url)
		url = strings.Replace(url, "!webp!webp", "!webp", 1)
	}
	return url
}
