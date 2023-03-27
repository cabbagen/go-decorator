package cache

import (
	"time"
	"strconv"
	"go-decorator/config"
	"github.com/go-redis/redis/v7"
)

type RedisCache struct {
	Type         string
	Client       *redis.Client
}

var defaultRedisCache *RedisCache

func NewRedisCache() *RedisCache {
	redisCache := RedisCache { "redis", nil }
	defaultRedisCache = &redisCache

	return &redisCache
}

func GetRedisCacheInstance() *RedisCache {
	if defaultRedisCache.Type != "" {
		return defaultRedisCache
	}
	return NewRedisCache()
}

func (rc *RedisCache) Connect() {
	db, _ := strconv.Atoi(config.CacheConfig["db"])
	addr, password := config.CacheConfig["addr"], config.CacheConfig["password"]

	rc.Client = redis.NewClient(&redis.Options { Addr: addr, Password: password, DB: db })
}

func (rc *RedisCache) Destroy() {
	rc.Client.Close()
}

func (rc *RedisCache) GetList(key string, start, stop int64) ([]string, error) {
	isExist, error := rc.Client.Exists(key).Result()

	if error != nil {
		return []string{}, error
	}
	if isExist == 0 {
		return []string{}, nil
	}
	return rc.Client.LRange(key, start, stop).Result()
}

func (rc *RedisCache) PushList(key string, values ...string) (int64, error) {
	return rc.Client.RPush(key, values).Result()
}

func (rc *RedisCache) UnShiftList(key string, values ...string) (int64, error) {
	return rc.Client.LPush(key, values).Result()
}

func (rc *RedisCache) GetSet(key string) ([]string, error) {
	return rc.Client.SMembers(key).Result()
}

func (rc *RedisCache) PushSet(key string, values ...string) (int64, error) {
	return rc.Client.SAdd(key, values).Result()
}

func (rc *RedisCache) HSet(key string, field string, values ...interface{}) (bool, error) {
	return rc.Client.HSet(key, field, values).Result()
}

func (rc *RedisCache) HGet(key string, field string) (string, error) {
	return rc.Client.HGet(key, field).Result()
}

func (rc *RedisCache) HDel(key string, fields ...string) (int64, error) {
	return rc.Client.HDel(key, fields...).Result()
}

func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return rc.Client.Set(key, value, expiration).Result()
}

func (rc *RedisCache) Get(key string) (string, error) {
	return rc.Client.Get(key).Result()
}

func (rc *RedisCache) Del(keys ...string) (int64, error) {
	return rc.Client.Del(keys...).Result()
}
