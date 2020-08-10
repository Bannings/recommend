package internal_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"recommend/golbal"
	"recommend/internal_api/proto"
	"recommend/log"
	"recommend/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	SectionParamMap = map[string]proto.Params{
		golbal.INDEX_PERSONAL:   {RuleKeys: []string{"hot"}, RuleSort: "ctr", RuleType: proto.RequestRuleType_HOT, Len: 6},             //猜你喜欢看
		golbal.INDEX_HIGHLY_NEW: {RuleKeys: []string{"hot"}, RuleSort: "recommend", RuleType: proto.RequestRuleType_HOT, Len: 4},       //入坑1秒
		golbal.INDEX_NEW:        {RuleKeys: []string{"hot"}, RuleSort: "new", RuleType: proto.RequestRuleType_HOT, Len: 6},             //X月新番
		golbal.INDEX_NOVEL:      {RuleKeys: []string{"小说漫改"}, RuleSort: "recommend", RuleType: proto.RequestRuleType_TOPIC, Len: 4},    //小说漫改
		golbal.INDEX_FREE_NEW:   {RuleKeys: []string{"free"}, RuleSort: "recommend", RuleType: proto.RequestRuleType_CLASSIFY, Len: 6}, //免费漫画
		golbal.INDEX_VIP:        {RuleKeys: []string{"vip"}, RuleSort: "recommend", RuleType: proto.RequestRuleType_CLASSIFY, Len: 4},  //vip,
	}
)

type recallResponse struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Data   RecallData `json:"data"`
}

type RecallData struct {
	Recall []string `json:"recall"`
}

func recallDataRequest(sUrl string) []string {
	request, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		log.Errorf("get sort requests failed, url:%s", sUrl)
		return nil
	}
	var recallResult recallResponse
	resp, err := reacllClient.Do(request)
	if err != nil {
		return nil
	}
	if resp != nil {
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}
			err = json.Unmarshal([]byte(respBytes), &recallResult)
			if err != nil {
				log.Errorf("get sort requests failed, url:%s", sUrl)
				return nil
			}
		}
	}
	return recallResult.Data.Recall
}

func getRecallData(uid int64, gender, url string, length ...int) []string {
	var lenRecall int
	if len(length) > 0 {
		lenRecall = length[0]
	} else {
		lenRecall = 60
	}

	sUrl := fmt.Sprintf(url+"?platform=samh&algo=all&uid=%d&sgender=%s&len=%d", uid, gender, lenRecall)
	return recallDataRequest(sUrl)
}

func getNewUserRecallData(uid int64, gender, url string, length ...int) []string {
	var lenRecall int
	if len(length) > 0 {
		lenRecall = length[0]
	} else {
		lenRecall = 60
	}

	sUrl := fmt.Sprintf(url+"?platform=samh&algo=all&uid=%d&sgender=%s&len=%d", uid, gender, lenRecall)
	return recallDataRequest(sUrl)
}

func getMixCandidates(uid int64, gender string) []string {
	conf := golbal.GetConfig()
	recallURL := conf.RuleRecallURL
	personalURL := conf.PersonalRecallURL
	urls := []string{recallURL, personalURL}
	var Candidates []string
	candidateChan := make(chan string, 130) //
	var waitGroup = new(sync.WaitGroup)
	waitGroup.Add(2)
	for _, url := range urls {
		go func(uid int64, gender, url string) {
			data := getRecallData(uid, gender, url)
			for _, comicId := range data {
				if comicId != "" {
					candidateChan <- comicId
				}
			}
			waitGroup.Done()
		}(uid, gender, url)
	}
	waitGroup.Wait()
	close(candidateChan)
	comicMap := make(map[string]int)
	for comicId := range candidateChan {
		//去重
		if _, ok := comicMap[comicId]; !ok {
			comicMap[comicId] = 0
			Candidates = append(Candidates, comicId)
		}
	}
	if len(Candidates) < 120 {
		ruleDate := model.GetStaticRuleData(model.OldUser).HighlyCandidate.MaleCandidates //通过候选集补全
		size := 120 - len(Candidates)
		if len(ruleDate) >= size {
			alternate := ruleDate[:size:size]
			Candidates = append(Candidates, alternate...)
		}
	}

	if len(Candidates) > 60 {
		return Candidates
	} else {
		return nil
	}
}

func getRuleCandidates(uid int64, gender string) []string {
	conf := golbal.GetConfig()
	recallURL := conf.RuleRecallURL
	Candidates := getNewUserRecallData(uid, gender, recallURL, 120)

	if len(Candidates) < 120 {
		ruleDate := model.GetStaticRuleData(model.OldUser).HighlyCandidate.MaleCandidates //通过候选集补全
		size := 120 - len(Candidates)
		if len(ruleDate) >= size {
			alternate := ruleDate[:size:size]
			Candidates = append(Candidates, alternate...)
		}
	}
	return Candidates
}

func GetCandidates(userInfo model.UserInfo) []string {
	uRead := userInfo.ReadCartoons
	readComics := strings.Split(uRead, "|")
	var candidates []string
	if userInfo.IsNewer {
		candidates = getRuleCandidates(userInfo.Uid, userInfo.SGender)
	} else {
		candidates = getMixCandidates(userInfo.Uid, userInfo.SGender)
	}
	return removeDuplicate(candidates, readComics)
}

func GetRelateRecommend(gender string, cid int) []string {
	conf := golbal.GetConfig()
	url := conf.RelateURL
	sUrl := fmt.Sprintf(url+"?platform=samh&algo=cf&cid=%d&sgender=%s&len=15", cid, gender)
	return recallDataRequest(sUrl)
}

func GetUserInterest(userInfo model.UserInfo, interest string, length ...int) []string {
	var recallLength int
	if len(length) > 0 {
		recallLength = length[0]
	} else {
		recallLength = 35
	}
	conf := golbal.GetConfig()
	url := conf.InterestURL
	var sUrl string
	if interest == "" {
		sUrl = fmt.Sprintf(url+"?platform=samh&algo=all&uid=%d&sgender=%s&len=%d", userInfo.Uid, userInfo.SGender, recallLength)
	} else {
		sUrl = fmt.Sprintf(url+"?platform=samh&algo=all&uid=%d&sgender=%s&interest=%s&len=%d", userInfo.Uid, userInfo.SGender, interest, recallLength)
	}

	return recallDataRequest(sUrl)
}

func removeDuplicate(candidates, readComics []string) []string {
	readComicMap := make(map[string]int)
	for _, comic := range readComics {
		if _, ok := readComicMap[comic]; !ok {
			readComicMap[comic] = 0
		}
	}
	var result []string
	for _, candidate := range candidates {
		if _, ok := readComicMap[candidate]; !ok {
			result = append(result, candidate)
		}
	}
	return result
}

func GetRecommendRecall(userInfo model.UserInfo, bussinessId string) (recommendResult *model.RecommendResult, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("no data return")
			return
		}
	}()
	sections := golbal.RecommendRecallColumns
	conn := GetMiddlewareRPC()
	client := proto.NewMiddlewareRPCClient(conn)
	var gender proto.Gender
	if userInfo.Gender == 0 {
		gender = proto.Gender_MALE
	} else {
		gender = proto.Gender_FEMALE
	}
	user := proto.UserInfo{Uid: userInfo.Uid, Udid: userInfo.UdId, Gender: gender}
	var params []*proto.Params
	for _, section := range sections {
		if param, ok := SectionParamMap[section]; ok {
			params = append(params, &param)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req := proto.MultiplyRecommendRequest{
		User:        &user,
		BussinessId: bussinessId,
		Params:      params,
	}
	resp, err := client.GetMultiplyRecommend(ctx, &req)
	if err != nil {
		log.Errorf("recall rpc error: %v", err)
		return nil, err
	}
	if resp.Data != nil {
		recommendMap := getRecommend(resp.Data, sections)
		//log.Infof("get recommendMap:%v", recommendMap)
		passthroughs := getPassthrough(resp.Data, bussinessId)
		if recommendMap != nil {
			recommendResult := model.RecommendResult{RecommendMap: recommendMap, ABTestExps: passthroughs}
			return &recommendResult, nil
		} else {
			return nil, fmt.Errorf("wrong length of multiply recommend data")
		}
	} else {
		return nil, fmt.Errorf("no data return")
	}
}

func GetRecommendRecallMore(userInfo model.UserInfo, bussinessId, sectionId string) (recommendResult []int32, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("no data return")
			return
		}
	}()
	conn := GetMiddlewareRPC()
	client := proto.NewMiddlewareRPCClient(conn)
	var gender proto.Gender
	if userInfo.Gender == 0 {
		gender = proto.Gender_MALE
	} else {
		gender = proto.Gender_FEMALE
	}
	user := proto.UserInfo{Uid: userInfo.Uid, Udid: userInfo.UdId, Gender: gender}
	var params []*proto.Params
	if param, ok := SectionParamMap[sectionId]; ok {
		params = []*proto.Params{&param}
	} else {
		return nil, fmt.Errorf("wrong section id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req := proto.MultiplyRecommendRequest{
		User:        &user,
		BussinessId: bussinessId,
		Params:      params,
	}
	resp, err := client.GetSectionMoreRecommend(ctx, &req)
	if err != nil {
		log.Errorf("recall rpc error: %v", err)
		return nil, err
	}
	if resp.Data != nil {
		return resp.Data.CartoonID, nil
	} else {
		return nil, fmt.Errorf("no data return")
	}
}

func GetRecommendRecallRenew(userInfo model.UserInfo, bussinessId, sectionId string) (recommendResult []int32, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("no data return")
			return
		}
	}()
	conn := GetMiddlewareRPC()
	client := proto.NewMiddlewareRPCClient(conn)
	var gender proto.Gender
	if userInfo.Gender == 0 {
		gender = proto.Gender_MALE
	} else {
		gender = proto.Gender_FEMALE
	}
	user := proto.UserInfo{Uid: userInfo.Uid, Udid: userInfo.UdId, Gender: gender}
	var params []*proto.Params
	if param, ok := SectionParamMap[sectionId]; ok {
		params = []*proto.Params{&param}
	} else {
		return nil, fmt.Errorf("wrong section id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	req := proto.MultiplyRecommendRequest{
		User:        &user,
		BussinessId: bussinessId,
		Params:      params,
	}
	resp, err := client.GetMultiplyRecommend(ctx, &req)
	if err != nil {
		log.Errorf("recall rpc error: %v", err)
		return nil, err
	}
	if resp.Data != nil {
		return resp.Data.CartoonIDs[0].Cids, nil
	} else {
		return nil, fmt.Errorf("no data return")
	}
}

func getRecommend(multiplyData *proto.MultiplyData, sections []string) map[string][]int32 {
	var recommendMap map[string][]int32
	recommendMap = make(map[string][]int32)
	if multiplyData == nil {
		return nil
	}
	if len(multiplyData.CartoonIDs) != len(sections) {
		log.Errorf("len mismatching,%d, %d", len(multiplyData.CartoonIDs), len(sections))
		return nil
	}
	for i, section := range sections {
		recommendMap[section] = multiplyData.CartoonIDs[i].Cids
		//log.Infof("get section:%s cids %v", section, multiplyData.CartoonIDs[i].Cids)
	}
	return recommendMap
}

func getPassthrough(multiplyData *proto.MultiplyData, bussinessId string) []model.ABExps {
	if !multiplyData.PassthroughActive {
		return nil
	}
	var passthroughs []model.ABExps
	for _, multiplyPassthrough := range multiplyData.Passthroughs {
		if multiplyPassthrough.Udid != "" {
			passthrough := model.ABExps{
				Udid:         multiplyPassthrough.Udid,
				Uid:          strconv.FormatInt(multiplyPassthrough.Uid, 10),
				ExpId:        multiplyPassthrough.ExpId,
				BucketId:     multiplyPassthrough.BucketId,
				BucketName:   multiplyPassthrough.BucketName,
				AlgId:        multiplyPassthrough.AlgId,
				RecServiceId: bussinessId,
				ServeTime:    multiplyPassthrough.ServeTime,
			}
			passthroughs = append(passthroughs, passthrough)
		}
	}
	return passthroughs
}
