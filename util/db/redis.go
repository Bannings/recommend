package db

import (
	"github.com/go-redis/redis"
	"recommend/golbal"
	"recommend/log"
	"strings"
	"time"
)

//TODO panic when client is nil

var (
	ComicInfoRedis  *RedisItems
	PersonalRead    *RedisItems //已读曝光
	PersonalInterst *RedisItems //feed流曝光
	Exposure        *RedisItems //首页固定栏曝光
	RedisCli        *redis.Client
	FeedRedisCli    *redis.Client
)

type RedisDb struct {
	RedisDataDb *redis.Client
	PoolSize    int
}

func NewRedisClient(conf golbal.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Host,
		Password:     conf.Password,
		PoolSize:     conf.PoolSize,
		MinIdleConns: 5,
		DialTimeout:  time.Duration(100) * time.Millisecond,
		ReadTimeout:  time.Duration(100) * time.Millisecond,
		WriteTimeout: time.Duration(100) * time.Millisecond,
	})
	return client
}

func NewRedisClientDB(conf golbal.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Host,
		Password:     conf.Password,
		PoolSize:     conf.PoolSize,
		DB:           conf.DB,
		MinIdleConns: 5,
		DialTimeout:  time.Duration(100) * time.Millisecond,
		ReadTimeout:  time.Duration(100) * time.Millisecond,
		WriteTimeout: time.Duration(100) * time.Millisecond,
	})
	return client
}

func RedisPGet(redisCli *redis.Client, keys []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := redisCli.Pipeline()
	for _, k := range keys {
		cmders = append(cmders, p.Get(k))
	}
	_, err := p.Exec()

	return cmders, err
}

func GetComicClassifyKeys(comicIds []string) []string {
	var result []string
	for _, comicId := range comicIds {
		key := "smh_comic_classify:" + comicId
		result = append(result, key)
	}
	return result
}

type RedisItems struct {
	RedisCli *redis.Client
	Prefix   string
	Expire   int64
	Len      int64
}

func (r *RedisItems) ItemGetKey(keys ...string) string {
	return r.Prefix + "_" + strings.Join(keys, "_")
}

func (r *RedisItems) ItemZAdd(values []string, key string) error {
	zmembers := make([]redis.Z, 0, len(values))
	for _, id := range values {
		zmembers = append(zmembers, redis.Z{Score: float64(time.Now().Unix()), Member: id})
	}
	p := r.RedisCli.Pipeline()
	expire := time.Duration(r.Expire) * time.Second
	p.Expire(key, expire)
	err := p.ZAdd(key, zmembers...).Err()

	cmdSetLen := p.ZCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err := r.RedisCli.ZRemRangeByRank(key, 0, setLen-r.Len-1).Err()
		if err != nil {
			log.Error(err)
		}
	}
	return err
}

func (r *RedisItems) ItemPGet(keys []string) ([]*redis.StringCmd, error) {
	var cmders []*redis.StringCmd
	p := r.RedisCli.Pipeline()
	for _, k := range keys {
		cmders = append(cmders, p.Get(r.ItemGetKey(k)))
	}
	_, err := p.Exec()

	return cmders, err
}

func (r *RedisItems) ItemSetSAdd(values []string, key string) error {

	p := r.RedisCli.Pipeline()
	err := p.SAdd(key, values).Err()
	expire := time.Duration(r.Expire) * time.Second
	p.Expire(key, expire)
	cmdSetLen := p.SCard(key)
	_, err = p.Exec()
	setLen := cmdSetLen.Val()
	if setLen > r.Len {
		err = r.RedisCli.SPopN(key, setLen-r.Len).Err()
		//errors_.CheckCommonErr(err)
	}
	return err
}

func InitRedis() {
	conf := golbal.GetConfig()
	RedisCli = NewRedisClient(conf.RedisCli)
	FeedRedisCli = NewRedisClientDB(conf.FeedRedisCli)
	PersonalRead = &RedisItems{RedisCli: RedisCli, Prefix: conf.RedisItems.Read.Prefix, Expire: conf.RedisItems.Read.Expire, Len: conf.RedisItems.Read.Len}
	PersonalInterst = &RedisItems{RedisCli: RedisCli, Prefix: conf.RedisItems.Interest.Prefix, Expire: conf.RedisItems.Interest.Expire, Len: conf.RedisItems.Interest.Len}
	Exposure = &RedisItems{RedisCli: RedisCli, Prefix: conf.RedisItems.Exposure.Prefix, Expire: conf.RedisItems.Exposure.Expire, Len: conf.RedisItems.Exposure.Len}
	ComicInfoRedis = &RedisItems{RedisCli: RedisCli, Prefix: conf.RedisItems.ComicInfoRedis.Prefix, Expire: conf.RedisItems.ComicInfoRedis.Expire, Len: conf.RedisItems.ComicInfoRedis.Len}
}
