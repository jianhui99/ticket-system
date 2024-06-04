package redis

import (
	"context"
	"fmt"
	"strconv"
	"ticket-system/config"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisLimiter *redis_rate.Limiter
var ctx = context.Background()

func Init() {

	addr := fmt.Sprintf("%s:%d", config.Conf.Database.RedisConf.Host, config.Conf.Database.RedisConf.Port)
	dbName := strconv.Itoa(config.Conf.Database.RedisConf.Db)
	db, _ := strconv.ParseUint(dbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   config.Conf.Database.RedisConf.Password,
		DB:         int(db),
		PoolSize:   config.Conf.Database.RedisConf.PoolSize,
		MaxRetries: 1,
	})

	_, connectErr := client.Ping(ctx).Result()

	if connectErr != nil {
		panic("Failed to connect to Redis")
	}

	flushErr := client.FlushDB(ctx).Err()

	if flushErr != nil {
		panic("Failed to flush Redis")
	}

	limiter := redis_rate.NewLimiter(client)

	RedisClient = client
	RedisLimiter = limiter

	fmt.Println("Redis init success")
}

func GetCache(key string) string {
	return RedisClient.Get(ctx, key).Val()
}

func SetCache(key string, value interface{}, expiration time.Duration) bool {
	isSetSuccesfully, err := RedisClient.SetNX(ctx, key, value, expiration).Result()

	if err != nil {
		fmt.Println("Error setting cache:", err)
		return false
	}

	return isSetSuccesfully
}

func DelCache(key string) {
	err := RedisClient.Del(ctx, key).Err()

	if err != nil {
		return
	}
}

func SetRListCache(key string, value interface{}, expiration time.Duration) int64 {
	isSetSuccesfully, err := RedisClient.RPush(ctx, key, value).Result()

	if err != nil {
		fmt.Println("Error setting list cache:", err)
		return 0
	}

	return isSetSuccesfully
}
