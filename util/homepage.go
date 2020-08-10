package util

import (
	"fmt"
	"github.com/xndm-recommend/go-utils/tools"
	"github.com/xndm-recommend/go-utils/type_"
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"recommend/util/recommend_api"
	"recommend/util/timed_task"
	"strconv"
	"strings"
)

//V3.0以下版本使用
func HomepageTopApplyV2(userInfo model.UserInfo) []model.SectionDataConfig {

	var recommendAll []string
	topColumns := getTopColumns(userInfo.IsNewer)
	candidates := getHomepageCandidates(userInfo)
	topCols := make([]model.SectionDataConfig, len(topColumns), len(topColumns)+golbal.ComicListNum+1)

	for index, sectionId := range topColumns {
		section, recommendList := getFixedSection(sectionId, userInfo, candidates)
		if section != nil {
			topCols[index] = *section
		}
		if len(recommendList) > 0 {
			for _, comicId := range recommendList {
				recommendAll = append(recommendAll, comicId)
			}
		}
	}
	if version2int(userInfo.AppVersion) >= 2101 || version2int(userInfo.AppVersion) <= 2108 {
		comicListSections, _ := internal_api.GetRecommendComicList(userInfo)
		if len(comicListSections) > 10 {
			comicListSections = comicListSections[:10]
		}
		temp := GetComicListSections(comicListSections, userInfo, 8)
		topCols = append(topCols, temp...)
	}

	// redis 数据更新
	updateUserReadRedis(userInfo, recommendAll)

	return topCols
}

//当RPC调用中间层获取推荐结果失败时，才会采用该方法，该方法为老版本获取推荐结果，通过内部restful API获取
func HomepageTopApplyV3(userInfo model.UserInfo) []model.SectionDataConfig {
	candidates := getHomepageCandidates(userInfo)
	var recommendAll, topColumns []string

	topColumns = getColumns(userInfo.AppVersion)
	var topCols []model.SectionDataConfig
	adInfos := internal_api.GetOperationInfo(userInfo, "HomeHorzBanner")
	for index, sectionId := range topColumns {
		adInfo := getAdinfo(adInfos, index)
		section, recommendList := getFixedSectionV2(sectionId, userInfo, candidates, adInfo)
		if section != nil {
			topCols = append(topCols, *section)
		}
		if len(recommendList) > 0 {
			for _, comicId := range recommendList {
				recommendAll = append(recommendAll, comicId)
			}
		}
	}
	// redis 数据更新
	updateUserReadRedis(userInfo, recommendAll)

	return topCols
}

func HomepageTopApplyV4(userInfo model.UserInfo) []model.SectionDataConfig {

	var recommendAll, topColumns []string
	topColumns = getColumns(userInfo.AppVersion)
	var topCols []model.SectionDataConfig
	adInfos := internal_api.GetOperationInfo(userInfo, "HomeHorzBanner")
	conf := golbal.GetConfig()
	recommendResult, err := internal_api.GetRecommendRecall(userInfo, conf.BusinessId)
	if err != nil {
		return HomepageTopApplyV3(userInfo)
	}
	for index, sectionId := range topColumns {
		adInfo := getAdinfo(adInfos, index)
		section, recommendList := getFixedSectionV3(sectionId, userInfo, *recommendResult, adInfo)
		if section != nil {
			topCols = append(topCols, *section)
		}
		if len(recommendList) > 0 {
			for _, comicId := range recommendList {
				recommendAll = append(recommendAll, comicId)
			}
		}
	}
	return topCols
}

func HomepageMore(userInfo model.UserInfo, sectionId string, pageNum, pageSize int) []model.SectionDataConfig {

	var recommendList []model.SectionDataConfig
	sectionIDInt, _ := strconv.Atoi(sectionId)

	switch {
	case sectionIDInt >= golbal.NEW_USER_COMIC_LIST:
		{
			outComicList := internal_api.GetNewUserComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	case sectionIDInt >= golbal.RECOMMEND_COMIC_LIST && sectionIDInt < golbal.NEW_USER_COMIC_LIST:
		{
			outComicList := internal_api.GetRecommendComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	default:
		{
			if pageNum >= 2 {
				data := model.SectionDataConfig{SectionId: sectionId}
				recommendList = []model.SectionDataConfig{data}
				return recommendList
			}
			conf := golbal.GetConfig()
			var recommendComics []string
			recommendResult, err := internal_api.GetRecommendRecallMore(userInfo, conf.BusinessId, sectionId)
			if len(recommendResult) == 0 || err != nil {
				recommendComics = recommend_api.GetTopRecommendMore(userInfo, sectionId, pageNum, pageSize)
			} else {
				recommendComics = int2String(recommendResult)
			}
			section := recommend_api.GetFixedColumnMoreSection(recommendComics, sectionId)
			recommendList = []model.SectionDataConfig{section}
		}
	}

	return recommendList
}

func HomepageRenew(userInfo model.UserInfo, sectionId string, pageNum, pageSize int) []model.SectionDataConfig {

	var recommendList []model.SectionDataConfig
	sectionIDInt, _ := strconv.Atoi(sectionId)

	switch {
	case sectionIDInt >= golbal.NEW_USER_COMIC_LIST:
		{
			outComicList := internal_api.GetNewUserComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	case sectionIDInt >= golbal.RECOMMEND_COMIC_LIST && sectionIDInt < golbal.NEW_USER_COMIC_LIST:
		{
			outComicList := internal_api.GetRecommendComicListMore(sectionId, pageNum, pageSize, userInfo)
			recommendList = GetComicListMoreSections(outComicList)
		}
	default:
		{
			conf := golbal.GetConfig()
			var recommendComics []string
			recommendResult, err := internal_api.GetRecommendRecallRenew(userInfo, conf.BusinessId, sectionId)
			if len(recommendResult) == 0 || err != nil {
				recommendComics = recommend_api.GetTopRecommendMore(userInfo, sectionId, pageNum, pageSize)
			} else {
				recommendComics = int2String(recommendResult)
			}
			section := recommend_api.GetFixedColumnMoreSection(recommendComics, sectionId)
			recommendList = []model.SectionDataConfig{section}
		}
	}

	return recommendList
}

func GetRelateRecommendSection(info model.UserInfo, cartoonName string, cid int, orderId int) *model.SectionDataConfig {
	if cartoonName == "" {
		comicInfo, err := recommend_api.QueryComicInfo(strconv.Itoa(cid))
		if err != nil {
			log.Error(err)
			return nil
		} else {
			cartoonName = comicInfo.ComicName
		}
	}
	section := model.RelateSection
	section.SectionName = fmt.Sprintf("看了《%s》的人还看过", cartoonName)
	section.SectionId = golbal.INDEX_RELATE
	section.SectionOrder = orderId + 1
	recommendList := internal_api.GetRelateRecommend(info.SGender, cid)
	recommendList = recommendRemoveDuplicate(recommendList, strconv.Itoa(cid))
	recommendList = tools.DifferenceStrLen(recommendList, info.Expose, 4)
	if recommendList != nil && len(recommendList) >= 4 {
		comicInfos := recommend_api.FetchComicInfos(db.ComicInfoRedis, recommendList)
		classfyMap := recommend_api.FetchComicsClassify(recommendList)
		for i, comicInfo := range comicInfos {
			if classfy, ok := classfyMap[comicInfo.ComicID]; ok {
				comicInfos[i].TypeList = classfy
			}
		}
		section.ComicInfo = comicInfos
		return section
	} else {
		return nil
	}

}

func GetInterestRecommendSection(userInfo model.UserInfo, cartoonName string, cid, orderId int) *model.SectionDataConfig {
	if cartoonName == "" {
		comicInfo, err := recommend_api.QueryComicInfo(strconv.Itoa(cid))
		if err != nil {
			log.Error(err)
			return nil
		} else {
			cartoonName = comicInfo.ComicName
		}
	}
	interestType := recommend_api.GetComicClassify(strconv.Itoa(cid))
	if interestType == "" {
		return nil
	}
	recommendList := internal_api.GetRelateRecommend(userInfo.SGender, cid)
	interest := internal_api.GetUserInterest(userInfo, interestType, 15)
	interest = append(interest, recommendList...)
	interest = tools.RmDuplicateStrLen(interest, len(interest))
	interest = recommendRemoveDuplicate(interest, strconv.Itoa(cid))
	interest = tools.DifferenceStrLen(interest, userInfo.Expose, 20)
	interest = recommend_api.GetStrListRandN(interest, 4)
	section := model.RelateSection
	section.SectionName = fmt.Sprintf("看了《%s》的人还看过", cartoonName)
	section.SectionId = golbal.INDEX_RELATE
	section.SectionOrder = orderId + 1
	if interest != nil && len(interest) >= 4 {
		comicInfos := recommend_api.FetchComicInfos(db.ComicInfoRedis, interest)
		classfyMap := recommend_api.FetchComicsClassify(interest)
		for i, comicInfo := range comicInfos {
			if classfy, ok := classfyMap[comicInfo.ComicID]; ok {
				comicInfos[i].TypeList = classfy
			}
		}
		section.ComicInfo = comicInfos
		return section
	} else {
		return nil
	}
}

func removeIPUdid(udId string) string {
	if strings.HasPrefix(udId, "ip") {
		return ""
	}
	return udId
}

func getFixedSection(sectionId string, userInfo model.UserInfo, candidates []string) (section *model.SectionDataConfig, recommendList []string) {

	//仅固定栏目顺序与新老用户有关，mini榜单，推荐书单等无关
	switch sectionId {
	case golbal.INDEX_BANNER:
		section = timed_task.GetBannerSection(userInfo, userInfo.DeviceType, sectionId)
	default:
		recommendList = recommend_api.GetTopRecommend(candidates, userInfo, sectionId)
		fixedSection := recommend_api.GetFixedColumnSection(recommendList, sectionId, userInfo.SGender, userInfo.IsNewer)
		section = &fixedSection
	}
	return section, recommendList
}

//V2与V3主要区别在于default中，v2通过老接口获取推荐结果为[]string,v3为新接口[]int32，调用V2，默认AB已失效，passthrough处理上有所不同
func getFixedSectionV2(sectionId string, userInfo model.UserInfo, candidates []string, adInfo *model.Operationinfo) (section *model.SectionDataConfig, recommendList []string) {
	conf := golbal.GetConfig()
	onlineService := model.OnlineService{RecServiceId: conf.BusinessId}
	switch sectionId {
	case golbal.INDEX_BANNER:
		section = timed_task.GetBannerSection(userInfo, golbal.BANNER_PLATFORM_ALL, sectionId)
		section.Passthrough.OnlineService = onlineService
	case golbal.INDEX_OPERATION_BANNER:
		operationinfo := internal_api.GetOperationInfo(userInfo, "HomeTOPBanner")
		section = recommend_api.GetOperationBannerSection(operationinfo, sectionId)
		section.Passthrough.OnlineService = onlineService
	case golbal.INDEX_OPERATION_AD:
		if adInfo == nil {
			return nil, recommendList
		}
		operationinfo := []model.Operationinfo{*adInfo}
		section = recommend_api.GetOperationBannerSection(operationinfo, sectionId)
		section.Passthrough.OnlineService = onlineService
	default:
		recommendList = recommend_api.GetTopRecommend(candidates, userInfo, sectionId)
		fixedSection, coverPassthrough := recommend_api.GetFixedColumnSectionV2(recommendList, sectionId, userInfo)
		if coverPassthrough != nil {
			fixedSection.Passthrough.ABTestExps = coverPassthrough
		} else {
			fixedSection.Passthrough.OnlineService = onlineService
		}
		section = &fixedSection
	}
	return section, recommendList
}

func getFixedSectionV3(sectionId string, userInfo model.UserInfo, recommendResult model.RecommendResult, adInfo *model.Operationinfo) (section *model.SectionDataConfig, recommendList []string) {
	conf := golbal.GetConfig()
	onlineService := model.OnlineService{RecServiceId: conf.BusinessId}
	switch sectionId {
	case golbal.INDEX_BANNER:
		section = timed_task.GetBannerSection(userInfo, golbal.BANNER_PLATFORM_ALL, sectionId)
		section.Passthrough.OnlineService = onlineService
	case golbal.INDEX_OPERATION_BANNER:
		operationinfo := internal_api.GetOperationInfo(userInfo, "HomeTOPBanner")
		section = recommend_api.GetOperationBannerSection(operationinfo, sectionId)
		section.Passthrough.OnlineService = onlineService
	case golbal.INDEX_OPERATION_AD:
		if adInfo == nil {
			//if version2int(userInfo.AppVersion) >= 3106 {
			//	section = recommend_api.GetADSection()
			//} else {
			return nil, recommendList
			//}
		} else {
			operationinfo := []model.Operationinfo{*adInfo}
			section = recommend_api.GetOperationBannerSection(operationinfo, sectionId)
			section.Passthrough.OnlineService = onlineService
		}
	default:
		recommendList = int2String(recommendResult.RecommendMap[sectionId])
		fixedSection, coverPassthrough := recommend_api.GetFixedColumnSectionV2(recommendList, sectionId, userInfo)
		if recommendResult.ABTestExps != nil {
			if coverPassthrough != nil {
				temp := recommendResult.ABTestExps
				temp = append(temp, coverPassthrough...)
				fixedSection.Passthrough.ABTestExps = temp
			} else {
				fixedSection.Passthrough.ABTestExps = recommendResult.ABTestExps
			}
		} else {
			fixedSection.Passthrough.OnlineService = onlineService
		}
		section = &fixedSection
	}
	return section, recommendList
}

func getAdinfo(infos []model.Operationinfo, index int) *model.Operationinfo {
	if infos == nil {
		return nil
	}
	//运营banner序号从1开始，序号2N为可展示banner或广告的位置
	for _, info := range infos {
		if info.OposOrder*2 == index {
			return &info
		}
	}
	return nil
}

func getHomepageCandidates(userInfo model.UserInfo) []string {
	recallList := internal_api.GetCandidates(userInfo)

	//candidates := internal_api.GetRecommendAlgoSort(recallList, userInfo, gobal.AlgoMap[gobal.INDEX_HIGHLY].Rule1)

	if len(recallList) > 0 {
		return recallList
	} else {
		return nil
	}
}

func updateUserReadRedis(userInfo model.UserInfo, recommendAll []string) {
	udid := removeIPUdid(userInfo.UdId)
	uidStr := strconv.FormatInt(userInfo.Uid, 10)
	key := db.PersonalRead.ItemGetKey(uidStr, udid, type_.IntToStr(userInfo.Gender), golbal.COMIC_SHOW)
	err := db.PersonalRead.ItemZAdd(recommendAll, key)
	if err != nil {
		log.Errorf("update user read redis error:%v", err)
	}
}

func getTopColumns(isNewUser bool) []string {
	if isNewUser {
		return golbal.NewColumns
	} else {
		return golbal.OldColumns
	}
}

func recommendRemoveDuplicate(inputs []string, cid string) []string {
	var outputs []string
	for _, input := range inputs {
		if input != cid {
			outputs = append(outputs, input)
		}
	}
	return outputs
}

func getColumns(version string) []string {
	if version2int(version) >= 3100 {
		return golbal.V3Columns
	} else {
		return golbal.V2Columns
	}
}

func int2String(cartoonIds []int32) []string {
	recommendList := make([]string, len(cartoonIds))

	for i, cartoonId := range cartoonIds {
		recommendList[i] = strconv.Itoa(int(cartoonId))
	}
	return recommendList
}

//version=2.2.9 转为int后为2209
func version2int(version string) int {
	if version == "" {
		return 0
	}
	versionInfo := strings.Split(string(version), ".")
	if len(versionInfo) < 3 {
		return 0
	}
	v, err := strconv.Atoi(versionInfo[0])
	if err != nil {
		return 0
	}
	s, err := strconv.Atoi(versionInfo[1])
	if err != nil {
		return 0
	}
	c, err := strconv.Atoi(versionInfo[2])
	if err != nil {
		return 0
	}
	intVersion := v*1000 + s*100 + c
	return intVersion
}
