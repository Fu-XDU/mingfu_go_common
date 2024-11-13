package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
	"mignfu_common/flags"
)

func NewConnOptionsFromFlags() *redis.Options {
	addr := fmt.Sprintf("%s:%d", flags.RedisHost, flags.RedisPort)
	return &redis.Options{
		Addr:         addr,
		Username:     flags.RedisUser,
		Password:     flags.RedisPasswd,
		DB:           flags.RedisDb,
		PoolSize:     flags.RedisPoolSize,
		MaxRetries:   3,
		MinIdleConns: flags.RedisPoolSize / 2,
	}
}

func Connect(options *redis.Options, initCallback func(*redis.Client) error) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(options)

	_, err = rdb.Ping().Result()
	if err != nil {
		log.Errorf("Connect to redis '%s'@'%s', DB:%d failed", options.Username, options.Addr, options.DB)
		return
	}

	if err = initCallback(rdb); err != nil {
		log.Errorf("Initialize mysql database '%s'@'%s', DB:%d failed, err: %v", options.Username, options.Addr, options.DB, err)
		return
	}

	log.Infof("Successfully connected to redis '%s'@'%s', DB:%d", options.Username, options.Addr, options.DB)
	return
}

func Subscribe(rdb *redis.Client, channels ...string) (ch <-chan *redis.Message, err error) {
	if len(channels) == 0 {
		return
	}
	pubSub := rdb.Subscribe(channels...)
	_, err = pubSub.Receive()
	if err != nil {
		log.Fatal(err)
	}
	ch = pubSub.Channel()
	return
}

func Get(rdb *redis.Client, key string) (val string, err error) {
	val, err = rdb.Get(key).Result()
	return
}

func Set(rdb *redis.Client, key, value string) (val string, err error) {
	val, err = rdb.Set(key, value, 0).Result()
	return
}

func Del(rdb *redis.Client, key string) (val int64, err error) {
	val, err = rdb.Del(key).Result()
	return
}
