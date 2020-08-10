package util

import (
	"fmt"
	"recommend/golbal"
	"recommend/model"
	"recommend/util/db"
	"recommend/util/recommend_api"
)

func SamhRecommendStreamApplyV2(userInfo model.UserInfo, pageSize int) []model.SectionDataConfig {
	sectionId := golbal.INDEX_TYPE
	recommendList := recommend_api.StreamRecommend(userInfo, pageSize)
	section := recommend_api.GetFixedColumnSection(recommendList, sectionId, userInfo.SGender, userInfo.IsNewer)
	return []model.SectionDataConfig{section}
}

func SamhRecommendStreamApplyV3(userInfo model.UserInfo, pageSize, pageNum int) []model.SectionDataConfig {

	recommendList := recommend_api.StreamRecommend(userInfo, pageSize)
	var section model.SectionDataConfig
	if pageNum == 1 {
		sectionId := golbal.OLD_INDEX_TYPE
		section = recommend_api.GetFixedColumnSection(recommendList, sectionId, userInfo.SGender, userInfo.IsNewer)
	} else {
		sectionId := golbal.INDEX_TYPE
		section = recommend_api.GetFixedColumnSection(recommendList, sectionId, userInfo.SGender, userInfo.IsNewer)
	}

	return []model.SectionDataConfig{section}
}

func SamhTypeApply(userInfo model.UserInfo, pageSize int, index string) []model.SectionDataConfig {
	sectionId := index
	recommendList := recommend_api.StreamRecommend(userInfo, pageSize)
	section := model.SectionDataConfig{}
	sectionName := golbal.ColumnNameMap[sectionId]
	if sectionId == golbal.INDEX_NEW {
		sectionName = fmt.Sprintf(sectionName, recommend_api.GetMonthName())
	}
	section.SectionName = sectionName
	section.ComicInfo = recommend_api.FetchComicInfos(db.ComicInfoRedis, recommendList)
	section.SectionId = sectionId

	return []model.SectionDataConfig{section}
}

func SamhApply(userInfo model.UserInfo, sectionId string) []model.SectionDataConfig {

	recommendList := recommend_api.GetRecommend(userInfo, sectionId)
	section := model.SectionDataConfig{}
	sectionName := golbal.ColumnNameMap[sectionId]
	if sectionId == golbal.INDEX_NEW {
		sectionName = fmt.Sprintf(sectionName, recommend_api.GetMonthName())
	}
	section.SectionName = sectionName
	section.ComicInfo = recommend_api.FetchComicInfos(db.ComicInfoRedis, recommendList)
	section.SectionId = sectionId

	return []model.SectionDataConfig{section}
}

func GetFeedStream(userInfo model.UserInfo, pageNum int) []model.SectionDataConfig {
	sectionId := golbal.INDEX_BEHAVIOR
	recommendList, recommendMap := recommend_api.FeedStreamRecommend(userInfo)
	//recommendList := internal_api.GetStreamRecommend(userInfo)

	if recommendList == nil {
		return []model.SectionDataConfig{}
	}
	section := recommend_api.GetFeedStreamSection(recommendList, sectionId, recommendMap)
	if pageNum != 1 {
		section.Config.ShowHeader = 0
	}
	return []model.SectionDataConfig{section}
}
