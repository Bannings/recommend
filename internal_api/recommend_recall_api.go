package internal_api

import (
	"encoding/json"
	"fmt"
	"github.com/xndm-recommend/go-utils/maths"
	"io/ioutil"
	"net/http"
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"time"
)

var (
	reacllClient = http.Client{Timeout: time.Duration(500) * time.Millisecond, Transport: t}
)

type recommendRecallReceive struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Data   RecallData `json:"data"`
}

// GetPersonalizedRecommendRecall 从rec_app根据参数algo获取个性化推荐结果
func GetPersonalizedRecommendRecall(userInfo model.UserInfo, length ...int) []string {
	if userInfo.Uid == 0 {
		return nil
	}
	var recallLength int
	if len(length) > 0 {
		recallLength = length[0]
	} else {
		recallLength = 12
	}
	conf := golbal.GetConfig()
	var algo string
	algo = "all"

	sUrl := fmt.Sprintf(conf.PersonalRecallURL+"?platform=samh&algo=%s&len=%d&uid=%d&sgender=%s", algo, recallLength, userInfo.Uid, userInfo.SGender)

	recommendRecall, err := GetRecall(sUrl)
	if err != nil {
		log.Error(err)
		return nil
	}
	return recommendRecall.Data.Recall
}

// GetRecommendResult 获取个性化推荐候选集结果
func GetRecommendResult(comicIds []string, userInfo model.UserInfo) (sortComicIds []string) {
	recommendResult := GetPersonalizedRecommendRecall(userInfo)
	if len(recommendResult) == 0 {
		return comicIds
	}
	return recommendResult
}

// GetRecommendResultSize 获取个性化推荐候选集结果，取size个漫画
func GetRecommendResultSize(comicIds []string, userInfo model.UserInfo, size int) (sortComicIds []string) {
	recommendList := GetPersonalizedRecommendRecall(userInfo, size)
	if len(recommendList) < size {
		length := maths.MinInt(len(comicIds), size)
		return comicIds[:length]
	}
	//recommendList = GetRecommendAlgoSort(recommendList, userInfo.Uid, userInfo.SGender)

	return recommendList
}

func GetRecall(url string) (*recommendRecallReceive, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var recallResult recommendRecallReceive
	resp, err := reacllClient.Do(request)
	if resp != nil {
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(respBytes), &recallResult)
			if err != nil {
				return nil, err
			}
		}
	}

	return &recallResult, err
}
