package recommend_api

import (
	"fmt"
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/internal_api/proto"
	"recommend/model"
	"recommend/util/db"
	"strconv"
	"time"
)

func GetFixedColumnSection(recommendList []string, sectionId, sGender string, isNewUser bool) model.SectionDataConfig {

	section := model.GetSection(sectionId)
	comicInfo := FetchComicInfos(db.ComicInfoRedis, recommendList)
	section.ComicInfo = comicInfo
	section.SectionId = sectionId
	if isNewUser {
		section.SectionOrder = golbal.NewColumnOrderMap[sectionId]
	} else {
		section.SectionOrder = golbal.OldColumnOrderMap[sectionId]
	}

	sectionName := golbal.ColumnNameMap[sectionId]
	if sectionId == golbal.INDEX_NEW {
		sectionName = fmt.Sprintf(sectionName, GetMonthName())
	}
	section.SectionName = sectionName
	return *section
}

func GetFixedColumnMoreSection(recommendResult []string, sectionId string) model.SectionDataConfig {
	section := model.GetSection(sectionId)
	comicInfo := FetchComicInfos(db.ComicInfoRedis, recommendResult)
	if sectionId == golbal.INDEX_HIGHLY_NEW {
		comicInfo = replaceDescription(comicInfo)
	}
	section.ComicInfo = comicInfo
	section.SectionId = sectionId

	sectionName := golbal.ColumnNameMap[sectionId]
	if sectionId == golbal.INDEX_NEW {
		sectionName = fmt.Sprintf(sectionName, GetMonthName())
	}
	section.SectionName = sectionName
	return *section
}

func GetFixedColumnSectionV2(recommendList []string, sectionId string, info model.UserInfo) (model.SectionDataConfig, []model.ABExps) {
	var section *model.SectionDataConfig
	var passthrough []model.ABExps
	section = model.GetSection(sectionId)
	comicInfo := FetchComicInfos(db.ComicInfoRedis, recommendList)
	if sectionId == golbal.INDEX_HIGHLY_NEW {
		comicInfo = replaceDescription(comicInfo)
	}
	if includeSpecialComic(recommendList) {
		coverData := internal_api.GetCoverRecommends(info)
		cover := getCover(coverData, recommendList)
		if cover != nil {
			passthrough = getPassthrough(cover, info)
			comicInfo = updateComicinfo(comicInfo, cover)

		}
	} else {
		passthrough = nil
	}
	section.ComicInfo = comicInfo
	section.SectionId = sectionId
	section.SectionOrder = golbal.V2ColumnOrderMap[sectionId]
	sectionName := golbal.ColumnNameMap[sectionId]
	if sectionId == golbal.INDEX_NEW {
		sectionName = fmt.Sprintf(sectionName, GetMonthName())
	}
	section.SectionName = sectionName
	return *section, passthrough
}

func GetOperationBannerSection(operationInfos []model.Operationinfo, sectionId string) *model.SectionDataConfig {
	var section *model.SectionDataConfig
	section = model.GetSection(sectionId)
	section.SectionId = sectionId
	section.SectionOrder = golbal.V3ColumnOrderMap[sectionId]
	section.SectionName = golbal.ColumnNameMap[sectionId]
	section.OperationInfo = operationInfos
	return section
}

func GetADSection() *model.SectionDataConfig {
	conf := golbal.GetConfig()
	onlineService := model.OnlineService{RecServiceId: conf.BusinessId}
	section := model.SectionDataConfig{DisplayType: 25, Passthrough: model.Passthrough{OnlineService: onlineService}} //DisplayType: 25广告专用样式
	return &section
}

func GetFeedStreamSection(recommendList []string, sectionId string, recommendMap map[string]string) model.SectionDataConfig {
	section := model.BehaviorSection
	section.SectionId = sectionId
	section.SectionName = golbal.ColumnNameMap[sectionId]
	comicInfos := FetchComicInfos(db.ComicInfoRedis, recommendList)
	classfyMap := FetchComicsClassify(recommendList)
	for i, comicInfo := range comicInfos {
		if classfy, ok := classfyMap[comicInfo.ComicID]; ok {
			comicInfos[i].TypeList = classfy
		}
		if based, ok := recommendMap[comicInfo.ComicID]; ok {
			cartoon, err := QueryComicInfo(based)
			if err == nil {
				comicInfos[i].Reason = fmt.Sprintf("和《%s》相似的作品", cartoon.ComicName)
			}
		}
	}
	section.ComicInfo = comicInfos
	return *section
}

func GetFeedStreamRelateSection(recommendList []string, sectionId, comicName string) model.SectionDataConfig {
	section := model.BehaviorRelateSection
	section.SectionId = sectionId
	section.SectionName = golbal.ColumnNameMap[sectionId]
	comicInfos := FetchComicInfos(db.ComicInfoRedis, recommendList)
	classfyMap := FetchComicsClassify(recommendList)
	for i, comicInfo := range comicInfos {
		if classfy, ok := classfyMap[comicInfo.ComicID]; ok {
			comicInfos[i].TypeList = classfy
			comicInfos[i].Reason = fmt.Sprintf("和《%s》相似的作品", comicName)
		}
	}
	section.ComicInfo = comicInfos
	return *section
}

func GetWxappColumnSection(recommendList []string, sectionId string) model.SectionDataConfig {

	section := model.GetSection(sectionId)
	section.ComicInfo = FetchComicInfos(db.ComicInfoRedis, recommendList)
	section.SectionId = sectionId
	section.SectionName = golbal.WxappColumnMap[sectionId]
	return *section
}

func GetComicListSection(recommendList []string, sectionId string, comicListSection model.OutComicList) model.SectionDataConfig {
	size := 4
	if i, ok := model.DisplayLenMap[model.GetSection(sectionId).DisplayType]; ok {
		size = i
	}
	if len(recommendList) >= size {
		recommendList = recommendList[:size:size]
	}
	section := model.GetSection(sectionId)
	section.SectionName = comicListSection.BookTitle
	section.SectionId = comicListSection.SectionID
	section.SectionSubtitle = comicListSection.Subtitle
	section.ComicInfo = FetchComicInfos(db.ComicInfoRedis, recommendList)
	return *section
}

func GetComicListMoreSection(recommendList []string, sectionId string, comicListSection model.OutComicList) model.SectionDataConfig {
	if len(recommendList) == 0 {
		recommendList = model.GetStaticRuleData(model.OldUser).HighlyCandidate.MaleCandidates
	}
	section := model.GetSection(sectionId)
	section.SectionName = comicListSection.BookTitle
	section.SectionId = comicListSection.SectionID
	section.SectionSubtitle = comicListSection.Subtitle
	section.ComicInfo = FetchComicInfos(db.ComicInfoRedis, recommendList)
	return *section
}

func GetMonthName() string {
	month := time.Now().Month()
	switch month {
	case time.January:
		return "一"
	case time.February:
		return "二"
	case time.March:
		return "三"
	case time.April:
		return "四"
	case time.May:
		return "五"
	case time.June:
		return "六"
	case time.July:
		return "七"
	case time.August:
		return "八"
	case time.September:
		return "九"
	case time.October:
		return "十"
	case time.November:
		return "十一"
	case time.December:
		return "十二"
	default:
		return "当"
	}
}

func includeSpecialComic(recommendList []string) bool {
	for _, comicId := range recommendList {
		if comicId == "200000" || comicId == "200001" || comicId == "108278" {
			return true
		}
	}
	return false
}

func updateComicinfo(comics []model.ComicInfo, covers []proto.MultiCoverDivision) []model.ComicInfo {
	for i := range comics {
		for _, cover := range covers {
			if comics[i].ComicID == strconv.Itoa(int(cover.ComicID)) {
				comics[i].OptionCover = cover.Cover
			}
		}

	}
	return comics
}

func getPassthrough(covers []proto.MultiCoverDivision, info model.UserInfo) []model.ABExps {
	if covers == nil {
		return nil
	}
	conf := golbal.GetConfig()
	result := make([]model.ABExps, len(covers))
	for i, cover := range covers {
		abExps := model.ABExps{
			ExpId:        cover.Passthrough.ExpId,
			Uid:          strconv.FormatInt(info.Uid, 10),
			Udid:         info.UdId,
			BucketId:     cover.Passthrough.BucketId,
			BucketName:   cover.Passthrough.BucketName,
			RecServiceId: conf.BusinessId,
			ServeTime:    time.Now().Unix(),
		}
		result[i] = abExps
	}
	return result
}

func getCover(coverRecommends []*proto.MultiCoverDivision, recommendList []string) []proto.MultiCoverDivision {
	var result []proto.MultiCoverDivision
	for _, cover := range coverRecommends {
		if cover.Cover != "" {
			for i := range recommendList {
				if strconv.Itoa(int(cover.ComicID)) == recommendList[i] && cover.Cover != "" {
					result = append(result, *cover)
				}
			}
		}
	}
	return result
}

func replaceDescription(comicInfos []model.ComicInfo) []model.ComicInfo {
	for i := range comicInfos {
		if description, ok := model.RecommendMaps.DescriptionMap[comicInfos[i].ComicID]; ok {
			comicInfos[i].ComicFeature = description
		}
	}
	return comicInfos
}
