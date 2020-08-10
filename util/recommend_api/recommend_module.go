package recommend_api

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/maths"
	"github.com/xndm-recommend/go-utils/tools"
	"math/rand"
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"strconv"
	"time"
)

func getHomepageCandidate(index string, gender int, isNewUser bool) []string {

	var userType model.UserType
	if isNewUser {
		userType = model.NewUser
	} else {
		userType = model.OldUser
	}
	var ruleMale, ruleFemale []string
	switch index {
	case golbal.INDEX_HIGHLY:
		ruleMale = model.GetStaticRuleData(userType).HighlyCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).HighlyCandidate.FemaleCandidates
	case golbal.INDEX_NEW:
		ruleMale = model.GetStaticRuleData(userType).NewCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).NewCandidate.FemaleCandidates
	case golbal.INDEX_VIP:
		ruleMale = model.GetStaticRuleData(userType).VipCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).VipCandidate.FemaleCandidates
	case golbal.INDEX_FREE:
		ruleMale = model.GetStaticRuleData(userType).FreeCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).FreeCandidate.FemaleCandidates
	case golbal.INDEX_PERSONAL:
		ruleMale = model.GetStaticRuleData(userType).PersonalCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).PersonalCandidate.FemaleCandidates
	case golbal.INDEX_FREE_NEW:
		ruleMale = model.GetStaticRuleData(userType).FreeCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).FreeCandidate.FemaleCandidates
	case golbal.INDEX_HIGHLY_NEW:
		ruleMale = model.GetStaticRuleData(userType).HighlyCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).HighlyCandidate.FemaleCandidates
	default:
		ruleMale = model.GetStaticRuleData(userType).HighlyCandidate.MaleCandidates
		ruleFemale = model.GetStaticRuleData(userType).HighlyCandidate.FemaleCandidates
	}

	if golbal.FEMALE_ID == gender {
		return ruleFemale
	} else if golbal.MIX_GENDER_ID == gender {
		return tools.MixStrList(ruleMale, ruleFemale)
	} else {
		return ruleMale
	}
}

func getHomepageCandidateWxapp(index string, gender int) []string {

	ruleRecommend := model.RuleRecommendWxapp
	var ruleMale, ruleFemale []string
	switch index {
	case golbal.INDEX_HIGHLY:
		ruleMale = ruleRecommend.HighlyCandidate.MaleCandidates
		ruleFemale = ruleRecommend.HighlyCandidate.FemaleCandidates
	case golbal.INDEX_NEW:
		ruleMale = ruleRecommend.NewCandidate.MaleCandidates
		ruleFemale = ruleRecommend.NewCandidate.FemaleCandidates
	case golbal.INDEX_PERSONAL:
		ruleMale = ruleRecommend.PersonalCandidate.MaleCandidates
		ruleFemale = ruleRecommend.PersonalCandidate.FemaleCandidates
	}
	if golbal.FEMALE_ID == gender {
		return ruleFemale
	} else if golbal.MIX_GENDER_ID == gender {
		return tools.MixStrList(ruleMale, ruleFemale)
	} else {
		return ruleMale
	}
}

func GetTopRecommend(Candidates []string, userInfo model.UserInfo, index string) (recommendList []string) {
	size := 6
	if i, ok := model.DisplayLenMap[model.GetSection(index).DisplayType]; ok {
		size = i
	}
	if len(Candidates) < size {
		if ruleTmp := getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer); len(ruleTmp) >= size {
			recommendList = ruleTmp[:size:size]
		}
	}

	rule := getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer) //获取当前固定栏所有候选集
	uidStr := strconv.FormatInt(userInfo.Uid, 10)
	key := db.Exposure.ItemGetKey(uidStr, userInfo.SGender)

	if index == golbal.INDEX_PERSONAL {
		candidates := internal_api.GetRecommendResultSize(rule, userInfo, 50)
		shows, err := db.RedisCli.ZRange(key, 0, -1).Result()
		if err != nil && err != redis.Nil {
			log.Errorf("get exposure err:%v", err)
		}
		recommendList = tools.DifferenceStrLen(candidates, shows, size)
		if len(recommendList) > 0 {
			if err = db.Exposure.ItemZAdd(recommendList, key); err != nil {
				log.Errorf("add exposure err:%v", err)
			}
		}
		if len(recommendList) >= size {
			return GetStrListRandN(recommendList, size)
		} else {
			if len(candidates) >= size {
				return GetStrListRandN(candidates, size)
			} else {
				return GetReturnResult(rule, size, true)
			}
		}
	} else {
		availiableCandidates := tools.IntersectStrList(Candidates, rule) //取出当前栏目可用集
		shows, err := db.RedisCli.ZRange(key, 0, -1).Result()
		if err != nil && err != redis.Nil {
			log.Errorf("get exposure err:%v", err)
		}

		availiableCandidates = tools.RmDuplicateStrLen(append(availiableCandidates, rule...), len(shows)+size) //去重
		recommendList = tools.DifferenceStrLen(availiableCandidates, shows, size)
		if len(recommendList) < size {
			if ruleTmp := getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer); len(ruleTmp) >= size {
				recommendList = ruleTmp[:size:size]
			}
		}
		if len(recommendList) != 0 {
			if err = db.Exposure.ItemZAdd(recommendList, key); err != nil {
				log.Errorf("add exposure err:%v", err)
			}
		}
	}
	return
}

func Random(inputs []string) (outputs []string) {
	defer func() {
		if e := recover(); e != nil {
			outputs = inputs
		}
	}()

	if len(inputs) <= 0 {
		return nil
	}
	for i := len(inputs) - 1; i > 0; i-- {
		rand.Seed(time.Now().Unix())
		num := rand.Intn(i + 1)
		if num != i {
			inputs[i], inputs[num] = inputs[num], inputs[i]
		}
	}
	outputs = inputs
	return outputs
}

func GetTopRecommendMore(userInfo model.UserInfo, index string, pageNum, pageSize int) (recommendList []string) {
	var rule []string

	rule = getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer)
	minSize := pageSize

	if index == golbal.INDEX_PERSONAL {
		recommendList = internal_api.GetRecommendResult(rule, userInfo)
		if len(recommendList) < minSize {
			recommendList = rule
		}
	} else {
		recommendList = rule
	}
	if pageNum == 0 {
		recommendList = GetStrListRandN(recommendList, pageSize)
	} else {
		recommendList, _ = tools.GetStrListNoLoop(recommendList, pageSize, pageNum)
	}
	return
}

func GetTopRecommendWxapp(userInfo model.UserInfo, index string) (recommendList []string) {
	var rule []string
	size := 4
	rule = getHomepageCandidateWxapp(index, userInfo.Gender)
	if index == golbal.INDEX_PERSONAL {
		recommendList = internal_api.GetRecommendResultSize(rule, userInfo, size)
	} else {
		if len(rule) >= size {
			recommendList = rule[:size:size]
		}
	}
	// 容错机制
	if len(recommendList) < size {
		ruleTmp := model.RuleRecommendOldUser.HighlyCandidate.MaleCandidates
		if len(ruleTmp) >= size {
			recommendList = ruleTmp[:size:size]
		}
	}
	return
}

func GetRecommend(userInfo model.UserInfo, index string) (recommendList []string) {

	size := 4
	if i, ok := model.DisplayLenMap[model.GetSection(index).DisplayType]; ok {
		size = i
	}

	var rule []string
	rule = getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer)
	if len(rule) >= size {
		recommendList = rule[:size:size]
	}

	if len(recommendList) < size {
		ruleTmp := getHomepageCandidate(index, userInfo.Gender, userInfo.IsNewer)
		if len(ruleTmp) >= size {
			recommendList = ruleTmp[:size:size]
		}
	}
	return
}

func StreamRecommend(userInfo model.UserInfo, pageSize int) (recommendList []string) {
	userId := strconv.FormatInt(userInfo.Uid, 10)
	key := db.PersonalInterst.ItemGetKey(userId, userInfo.SGender)
	showed, err := db.PersonalInterst.RedisCli.ZRange(key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		log.Errorf("get show err:%v", err)
	}
	var seed []string
	if userInfo.SGender == golbal.MALE_STR {
		seed = model.FeedSeed.GetMaleCandidates()
	} else {
		seed = model.FeedSeed.GetFemaleCandidates()
	}
	candidates := getAvailableCandidate(seed, showed, pageSize)
	if len(candidates) > 0 {
		err := db.PersonalInterst.ItemZAdd(candidates, key)
		if err != nil {
			log.Error(err)
		}
		return candidates
	} else {
		return GetReturnResult(seed, pageSize, false)
	}
}

func GetStrListRandN(s []string, size int) []string {
	cs := append(s[:0:0], s...)
	rand.Shuffle(len(cs), func(i, j int) {
		cs[i], cs[j] = cs[j], cs[i]
	})
	if size <= 0 || size >= len(s) {
		return cs
	}
	return cs[:size]
}

func FeedStreamRecommend(userInfo model.UserInfo) ([]string, map[string]string) {
	userId := strconv.FormatInt(userInfo.Uid, 10)
	recommendMap := make(map[string]string)
	key := db.PersonalInterst.ItemGetKey(userId, userInfo.SGender)
	showed, err := db.PersonalInterst.RedisCli.ZRange(key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		log.Errorf("get show err:%v", err)
	}
	var seed []string
	if userInfo.SGender == golbal.MALE_STR {
		seed = model.FeedSeed.GetMaleCandidates()
	} else {
		seed = model.FeedSeed.GetFemaleCandidates()
	}
	candidates := getAvailableCandidate(seed, showed, 5)
	behaviorRecommend, basedComic := getBehaviorRecommend(userInfo.Uid, userInfo.SGender)
	if behaviorRecommend != "" {
		candidates = append(candidates, behaviorRecommend)
	}
	if len(candidates) >= 5 {
		err := db.PersonalInterst.ItemZAdd(candidates, key)
		if err != nil {
			log.Error(err)
		}
		if basedComic != "" {
			recommendMap[behaviorRecommend] = basedComic
			return GetReturnResult(candidates, len(candidates), false), recommendMap
		} else {
			return GetStrListRandN(candidates, len(candidates)), nil
		}
	} else {
		return GetReturnResult(seed, 5, true), nil
	}
}

func getBehaviorRecommend(uid int64, sGender string) (recommendComic, basedComic string) {
	recommendKey := fmt.Sprintf("user_rt_event_cache::samh::ranking_list::click::%d", uid)

	based, err := db.FeedRedisCli.Get(recommendKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ""
		}
		log.Errorf("get behavior recommend fail %v", err)
		return "", ""
	}
	basedCid, err := strconv.Atoi(based)
	if err != nil {
		log.Errorf("wrong comic %s, err:%v", based, err)
		return "", ""
	}
	relateRecommend := internal_api.GetRelateRecommend(sGender, basedCid)
	_, err = db.FeedRedisCli.Del(recommendKey).Result()
	if err != nil {
		log.Errorf("Del behavior key fail %v", err)
	}
	if len(relateRecommend) > 0 {
		return relateRecommend[0], based
	} else {
		return "", ""
	}
}

func getAvailableCandidate(feedSeed, shows []string, size int) []string {

	candidates := tools.DifferenceStrLen(feedSeed, shows, size) //求差集
	return candidates
}

func GetReturnResult(candidates []string, size int, isSufficient bool) []string {
	if len(candidates) >= size {
		return GetStrListRandN(candidates, size)
	}
	if isSufficient {
		result := model.RuleRecommendOldUser.HighlyCandidate.MaleCandidates
		return GetStrListRandN(result, len(result))
	} else {
		if len(candidates) > 0 {
			return candidates
		} else {
			result := model.RuleRecommendOldUser.HighlyCandidate.MaleCandidates
			return result[maths.MinInt(size, len(result)):maths.MinInt(size, len(result))]
		}
	}
}
