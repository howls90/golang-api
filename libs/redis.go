package libs

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"

	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RedisSet("posts", []byte("[]"))
}

func RedisGet(key string) string {
	val, err := RDB.Get(ctx, key).Result()
	if err != nil {
		log.Error(err)
	}

	return val
}

func RedisSet(key string, value []byte) {
	if err := RDB.Set(ctx, key, value, 0).Err(); err != nil {
		log.Error(err)
	}
}
