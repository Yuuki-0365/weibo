package dao

import (
	"SmallRedBook/cache"
	"context"
	"github.com/gomodule/redigo/redis"
)

type LikeDao struct {
	redis.Conn
}

func NewLikeDao(ctx context.Context) *LikeDao {
	return &LikeDao{cache.RedisPool.Get()}
}
func ReadDataFromRedis() {
}
