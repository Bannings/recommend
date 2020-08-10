package internal_api

import (
	"encoding/json"
	"github.com/xndm-recommend/go-utils/tools"
	"io/ioutil"
	"net/http"
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"strconv"
	"time"
)

var (
	comicListClient = http.Client{Timeout: time.Duration(1) * time.Second, Transport: t}
)

type newUserComicListRequest struct {
	Uid          string `url:"uid"`
	DeviceId     string `url:"device_id"`
	UdId         string `url:"udid"`
	DeviceBrand  string `url:"device_brand"`
	DeviceModel  string `url:"device_model"`
	ScreenWidth  string `url:"screen_width"`
	ScreenHeight string `url:"screen_height"`
	Carrier      string `url:"carrier"`
	Access       string `url:"access"`
	Os           string `url:"os"`
	OsVersion    string `url:"os_version"`
	Channel      string `url:"channel"`
	AppVersion   string `url:"app_version"`
	UAddr        string `url:"addr"`
	Sgender      string `url:"sgender"`
	SectionID    int    `url:"section_id"`
	PageNum      string `url:"page_num"`
	Newer        int    `url:"newer"`
}

func (r *newUserComicListRequest) fromUserInfo(userInfo model.UserInfo) {
	r.Uid = strconv.FormatInt(userInfo.Uid, 10)
	r.UdId = userInfo.UdId
	r.DeviceId = userInfo.UdId
	r.DeviceBrand = userInfo.DeviceBrand
	r.DeviceModel = userInfo.DeviceModel
	r.ScreenWidth = userInfo.ScreenWidth
	r.ScreenHeight = userInfo.ScreenHeight
	r.Carrier = userInfo.Carrier
	r.Access = userInfo.Access
	r.Os = userInfo.Os
	r.OsVersion = userInfo.OsVersion
	r.Channel = userInfo.Channel
	r.AppVersion = userInfo.AppVersion
	r.UAddr = userInfo.UAddr
	r.PageNum = userInfo.PageNum
	if userInfo.IsNewer {
		r.Newer = 1
	} else {
		r.Newer = 0
	}
}

type responseComicList struct {
	Title      string   `json:"title"`
	SectionID  int      `json:"section_id"`
	CartoonIDs []string `json:"cartoon_ids"`
}

type responseJson struct {
	Msg  string              `json:"msg"`
	Flag int                 `json:"flag"`
	Data []responseComicList `json:"data"`
}

func toOutComicLists(rcl []responseComicList) []model.OutComicList {
	outComicLists := make([]model.OutComicList, 0, len(rcl))
	for _, resComicList := range rcl {
		outComicLists = append(outComicLists, model.OutComicList{BookTitle: resComicList.Title,
			CartoonIDs: resComicList.CartoonIDs, SectionID: strconv.Itoa(resComicList.SectionID), CartoonTypes: []string{""}})
	}
	return outComicLists
}

func GetNewUserComicList(userInfo model.UserInfo) (outComicLists []model.OutComicList, flag int) {

	var requestParam newUserComicListRequest
	requestParam.fromUserInfo(userInfo)
	if userInfo.Gender == golbal.FEMALE_ID {
		requestParam.Sgender = golbal.FEMALE_STR
	} else {
		requestParam.Sgender = golbal.MALE_STR
	}

	conf := golbal.GetConfig()
	sUrl := Struct2Url(requestParam, conf.NewUserBookListURL)

	response, err := GetComicList(sUrl)
	if err != nil {
		log.Error(err)
	}
	if response == nil {
		return []model.OutComicList{}, 0
	}
	outComicLists = toOutComicLists(response.Data)

	return outComicLists, response.Flag
}

func GetNewUserComicListMore(sectionID string, pageNum, pageSize int, userInfo model.UserInfo) (outComicList model.OutComicList) {
	var requestParam newUserComicListRequest
	requestParam.fromUserInfo(userInfo)

	if userInfo.Gender == golbal.FEMALE_ID {
		requestParam.Sgender = golbal.FEMALE_STR
	} else {
		requestParam.Sgender = golbal.MALE_STR
	}

	requestParam.SectionID, _ = strconv.Atoi(sectionID)
	conf := golbal.GetConfig()
	sUrl := Struct2Url(requestParam, conf.NewUserBookListMoreURL)
	response, err := GetComicList(sUrl)
	if err != nil {
		log.Error(err)
		return
	}
	outComicLists := toOutComicLists(response.Data)
	if len(outComicLists) > 0 {
		outComicList = outComicLists[0]
	} else {
		rule := model.RuleRecommendOldUser.HighlyCandidate.FemaleCandidates
		comicList := model.OutComicList{BookTitle: "", SectionID: sectionID, CartoonIDs: rule}
		outComicList = comicList
	}
	if pageNum == 0 {
		outComicList.CartoonIDs = tools.GetStrListRandN(outComicList.CartoonIDs, pageSize)
	} else {
		outComicList.CartoonIDs, _ = tools.GetStrListNoLoop(outComicList.CartoonIDs, pageSize, pageNum)
	}
	return outComicList
}

func GetComicList(url string) (*responseJson, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var resJson responseJson

	resp, err := comicListClient.Do(request)
	if resp != nil {
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(respBytes), &resJson)
			if err != nil {
				return nil, err
			}
		}
	}

	return &resJson, err
}

func GetRecommendComicList(userInfo model.UserInfo) (outComicLists []model.OutComicList, flag int) {
	var requestParam newUserComicListRequest
	requestParam.fromUserInfo(userInfo)
	if userInfo.Gender == golbal.FEMALE_ID {
		requestParam.Sgender = golbal.FEMALE_STR
	} else {
		requestParam.Sgender = golbal.MALE_STR
	}

	conf := golbal.GetConfig()
	recommendBookListURL := conf.RecommendBookListURL
	sUrl := Struct2Url(requestParam, recommendBookListURL)

	response, err := GetComicList(sUrl)
	if err != nil {
		log.Error(err)
	}
	if response == nil {
		return []model.OutComicList{}, 0
	}
	outComicLists = toOutComicLists(response.Data)
	return outComicLists, response.Flag
}

func GetRecommendComicListMore(sectionID string, pageNum, pageSize int, userInfo model.UserInfo) (outComicList model.OutComicList) {
	var requestParam newUserComicListRequest
	requestParam.fromUserInfo(userInfo)

	if userInfo.Gender == golbal.FEMALE_ID {
		requestParam.Sgender = golbal.FEMALE_STR
	} else {
		requestParam.Sgender = golbal.MALE_STR
	}

	requestParam.SectionID, _ = strconv.Atoi(sectionID)
	conf := golbal.GetConfig()
	sUrl := Struct2Url(requestParam, conf.RecommendBookListMoreURL)
	response, err := GetComicList(sUrl)
	if err != nil {
		log.Error(err)
	}
	outComicLists := toOutComicLists(response.Data)
	if len(outComicLists) > 0 {
		outComicList = outComicLists[0]
		outComicList.CartoonIDs = tools.GetStrListRandN(outComicList.CartoonIDs, 6)
	} else {
		rule := model.RuleRecommendOldUser.HighlyCandidate.MaleCandidates
		comicList := model.OutComicList{BookTitle: "", SectionID: sectionID, CartoonIDs: rule}
		outComicList = comicList
	}

	return outComicList
}
