package service

import (
	"SmallRedBook/dao"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

type LikeService struct {
	NoteId string `json:"note_id" form:"note_id"`
}

func (service *LikeService) LikeNote(ctx context.Context, userId string) tool.Response {
	key := userId + ":like:" + service.NoteId
	county := userId + ":like:count"
	likeDao := dao.NewLikeDao(ctx)
	defer likeDao.Close()
	script, err := os.ReadFile("service/like.lua")
	if err != nil {
		e.ThrowError(e.Error)
	}
	lua := redis.NewScript(2, string(script))
	res, err := lua.Do(likeDao.Conn, key, county)
	if err != nil {
		e.ThrowError(e.Error)
	}
	if res.(int64) == 0 {
		e.ThrowError(e.Error)
	} else if res.(int64) == 1 || res.(int64) == 2 {
		return e.ThrowSuccess()
	}
	return e.ThrowError(e.Error)
}

func (service *LikeService) GetLikeCount(ctx context.Context, userId string) tool.Response {
	key := userId + ":like:count"
	likeDao := dao.NewLikeDao(ctx)
	value, err := likeDao.Do("GET", key)
	if err != nil {
		e.ThrowError(e.ErrorRedis)
	}
	data := fmt.Sprintf("%s", value)
	return tool.Response{
		Status: e.Success,
		Msg:    e.GetMsg(e.Success),
		Data:   data,
	}
}
