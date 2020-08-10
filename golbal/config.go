package golbal

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var (
	conf       *Config
	configLock = new(sync.RWMutex)
)

type Config struct {
	DbFeatures               DBConfig    `json:"db_features"`
	DbRecall                 DBConfig    `json:"db_recall"`
	DbBanner                 DBConfig    `json:"db_banner"`
	DBComic                  DBConfig    `json:"db_comic"`
	RedisCli                 RedisConfig `json:"redis_cli"`
	FeedRedisCli             RedisConfig `json:"feed_redis_cli"`
	RedisItems               RedisItems  `json:"redis_items"`
	Kafka                    []string    `json:"kafka"`
	UserPortraitURL          string      `json:"user_portrait_url"`
	NewUserBookListURL       string      `json:"new_user_comic_list_url"`
	NewUserBookListMoreURL   string      `json:"new_user_comic_list_more_url"`
	RecommendBookListURL     string      `json:"recommend_book_list_url"`
	RecommendBookListMoreURL string      `json:"recommend_book_list_more_url"`
	OperationUrl             string      `json:"operation_url"`
	FeedStreamUrl            string      `json:"feed_stream_url"`
	RuleRecallURL            string      `json:"rule_recall_url"`
	PersonalRecallURL        string      `json:"personal_recall_url"`
	InterestURL              string      `json:"interest_url"`
	RelateURL                string      `json:"relate_url"`
	MiddlewareRPC            string      `json:"middleware_rpc"`
	CoverRPC                 string      `json:"cover_rpc"`
	BusinessId               string      `json:"business_id"`
	Port                     int         `json:"port"`
}

type DBConfig struct {
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

type RedisItems struct {
	ComicInfoRedis RedisItem `json:"Comic_info"`
	Read           RedisItem `json:"uid_personal_read"`
	Interest       RedisItem `json:"uid_personal_interest"`
	Exposure       RedisItem `json:"exposure"`
}

type RedisItem struct {
	Prefix string `json:"prefix"`
	Expire int64  `json:"expire"`
	Len    int64  `json:"len"`
}

func LoadConfig(jsonFile string) (*Config, error) {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	tmp := new(Config)
	if err = json.Unmarshal(file, tmp); err != nil {
		return nil, err
	}

	configLock.Lock()
	defer configLock.Unlock()
	conf = tmp
	return conf, nil
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return conf
}
