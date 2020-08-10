package recommend_api

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/xndm-recommend/go-utils/sets"
	"recommend/log"
	"recommend/model"
	"recommend/util/db"
	"strconv"
	"strings"
	"time"
)

func getNewTypeList(_type string) string {
	var tmp []interface{}
	var typeStr []string
	var MapNewType map[interface{}]interface{}
	for _, t := range strings.Split(_type, "|") {
		if newType, ok := MapNewType[t]; ok {
			tmp = append(tmp, newType)
		} else {
			tmp = append(tmp, t)
		}
	}
	hashSet := sets.NewSetFromSlice(tmp)
	typeSli := hashSet.SetToSlice()
	for _, t := range typeSli {
		typeStr = append(typeStr, t.(string))
	}
	return strings.Join(typeStr, "|")
}

func QueryComicInfo(comicID string) (model.ComicInfo, error) {

	var comicInfo model.ComicInfo
	err := db.ComicDataDB().Model(&model.ComicInfo{}).Where("cartoon_id=?", comicID).Find(&comicInfo).Error

	comicInfo.TypeList = getNewTypeList(comicInfo.TypeList)
	return comicInfo, err
}

func getCartoonClassify(cartoonId string) string {
	var cartoonType model.CartoonClassify
	err := db.UserFeatureDB().Model(&model.CartoonClassify{}).Where("cartoon_id=?", cartoonId).Scan(&cartoonType).Error
	if err != nil {
		log.Errorf("get cartoon type fail:%v,cartoon id:%s", err, cartoonId)
	}
	return cartoonType.MainType
}

func GetComicInfoMap(comicIDs []string) map[string]model.ComicInfo {
	comicInfoMap := make(map[string]model.ComicInfo, len(comicIDs))
	for _, id := range comicIDs {
		comicInfo, err := QueryComicInfo(id)
		//errors_.CheckErrSendEmail(errors.WithMessage(err, "comic_id:"+id))
		if err == nil {
			comicInfoMap[id] = comicInfo
		}
	}
	return comicInfoMap
}

func CacheComicInfo(redisInfo *db.RedisItems, comicInfo *model.ComicInfo) {

	comicBytes, err := json.Marshal(comicInfo)
	key := redisInfo.ItemGetKey(comicInfo.ComicID)
	expire := time.Duration(redisInfo.Expire) * time.Second
	err = redisInfo.RedisCli.Set(key, comicBytes, expire).Err()
	if err != nil {
		log.Error(err)
	}
}

func FetchComicInfos(comicItem *db.RedisItems, comicIDs []string) []model.ComicInfo {
	comicInfos := make([]model.ComicInfo, 0, len(comicIDs))
	cmders, err := comicItem.ItemPGet(comicIDs)
	if err != nil && err != redis.Nil {
		log.Error(err)
	}
	for i, cmd := range cmders {
		var tmpComicInfo model.ComicInfo
		postByte, err := cmd.Bytes()
		if nil == err {
			err = json.Unmarshal(postByte, &tmpComicInfo)
		}
		if nil != err {
			if comicIDs[i] == "" {
				continue
			}
			tmpComicInfo, err = QueryComicInfo(comicIDs[i])
			if err != nil {
				log.Errorf("query comic error:%v, comic_id:%v", err, comicIDs[i])
				continue
			} else {
				CacheComicInfo(comicItem, &tmpComicInfo)
			}
		}
		comicID, _ := strconv.Atoi(tmpComicInfo.ComicID)
		if tagInfo, ok := model.RecommendMaps.CartoonTagMap[comicID]; ok {
			tmpComicInfo.ComicRank = tagInfo.Tag
			tmpComicInfo.ComicCategory = tagInfo.Category
		}
		if _, ok := model.RecommendMaps.ComicFreeMap[tmpComicInfo.ComicID]; ok {
			tmpComicInfo.ComicTips = "本周限免"
		}
		comicInfos = append(comicInfos, tmpComicInfo)
	}
	return comicInfos
}

func FetchComicsClassify(comicIDs []string) map[string]string {
	classifyMap := make(map[string]string)
	keys := db.GetComicClassifyKeys(comicIDs)
	cmders, err := db.RedisPGet(db.RedisCli, keys)
	if err != nil && err != redis.Nil {
		log.Error(err)
	}
	for i, cmd := range cmders {
		classifyMap[comicIDs[i]], err = cmd.Result()
		if err != nil {
			classify := getCartoonClassify(comicIDs[i])
			classifyMap[comicIDs[i]] = classify
			CacheComicClassify(comicIDs[i], classify)
		}
	}
	return classifyMap
}

func GetComicClassify(comicID string) (classify string) {
	key := "smh_comic_classify:" + comicID
	classify, err := db.FeedRedisCli.Get(key).Result()
	if err != nil {
		classify = getCartoonClassify(comicID)
	}
	return classify
}

func CacheComicClassify(key, value string) {
	cacheKey := "smh_comic_classify:" + key
	statusCmd := db.RedisCli.Set(cacheKey, value, 86400*time.Second)
	if statusCmd.Err() != nil {
		log.Errorf("cache comic classify fail:%v", statusCmd.Err())
	}
}
