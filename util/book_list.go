package util

import (
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/model"
	"recommend/util/recommend_api"
)

func GetBookList(userInfo model.UserInfo) ([]model.SectionDataConfig, int) {
	comicLists, flag := internal_api.GetNewUserComicList(userInfo)
	if comicLists != nil {
		return GetComicListSections(comicLists, userInfo, 8), flag
	} else {
		return nil, 0
	}
}

func GetBookListV2(userInfo model.UserInfo) ([]model.SectionDataConfig, int) {
	var comicListSections []model.OutComicList
	var flag, order int

	comicListSections, flag = internal_api.GetRecommendComicList(userInfo)
	if len(comicListSections) > 10 {
		comicListSections = comicListSections[:10]
	}
	if comicListSections != nil {
		return GetRecommendComicListSections(comicListSections, order), flag
	} else {
		return nil, 0
	}
}

func GetComicListSections(comicListSections []model.OutComicList, userInfo model.UserInfo, order int) []model.SectionDataConfig {
	//首页只显示最开始10个书单
	length := 10
	if len(comicListSections) > 0 && len(comicListSections) < 10 {
		length = len(comicListSections)
	} else if len(comicListSections) == 0 {
		return []model.SectionDataConfig{}
	} else {
		comicListSections = comicListSections[:10]
	}

	comicSections := make([]model.SectionDataConfig, 0, length)
	for i, comicListSection := range comicListSections {
		var recommendList []string

		if len(comicListSection.CartoonIDs) >= 4 {
			recommendList = comicListSection.CartoonIDs[:4]
		} else {
			recommendList = comicListSection.CartoonIDs
		}
		var displayType string //书单样式切换
		if i%2 == 1 {
			displayType = golbal.INDEX_COMICLIST1
		} else {
			displayType = golbal.INDEX_COMICLIST2
		}
		sectionDataConfig := recommend_api.GetComicListSection(recommendList, displayType, comicListSection)
		sectionDataConfig.SectionOrder = i + order
		comicSections = append(comicSections, sectionDataConfig)
	}
	return comicSections
}

func GetComicListMoreSections(comicListSection model.OutComicList) []model.SectionDataConfig {
	var recommendList []string
	recommendList = comicListSection.CartoonIDs
	sectionDataConfig := recommend_api.GetComicListMoreSection(recommendList, comicListSection.SectionID, comicListSection)

	return []model.SectionDataConfig{sectionDataConfig}
}

func GetRecommendComicListSections(comicListSections []model.OutComicList, order int) []model.SectionDataConfig {
	//首页只显示最开始10个书单
	length := 10
	if len(comicListSections) > 0 && len(comicListSections) < 10 {
		length = len(comicListSections)
	} else if len(comicListSections) == 0 {
		return []model.SectionDataConfig{}
	} else {
		comicListSections = comicListSections[:10]
	}
	conf := golbal.GetConfig()
	comicSections := make([]model.SectionDataConfig, 0, length)
	for i, comicListSection := range comicListSections {
		var recommendList []string

		if len(comicListSection.CartoonIDs) >= 8 {
			recommendList = comicListSection.CartoonIDs[:8]
		} else {
			recommendList = comicListSection.CartoonIDs
		}
		var displayType string //书单样式切换
		if i%4 == 2 || i%4 == 0 {
			displayType = golbal.INDEX_BOOKLIST1
		} else if i%4 == 1 {
			displayType = golbal.INDEX_BOOKLIST2
		} else {
			displayType = golbal.INDEX_BOOKLIST3
		}
		sectionDataConfig := recommend_api.GetComicListSection(recommendList, displayType, comicListSection)
		sectionDataConfig.SectionOrder = i + order
		sectionDataConfig.Passthrough.OnlineService = model.OnlineService{RecServiceId: conf.BusinessId}
		comicSections = append(comicSections, sectionDataConfig)
	}
	return comicSections
}
