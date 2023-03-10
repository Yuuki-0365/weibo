package service

import (
	"SmallRedBook/dao"
	"SmallRedBook/model"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strings"
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
	defer likeDao.Close()
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

func (service *LikeService) ReadDataFromRedis() {
	likeKey := "*:like:*"
	likeCountKey := "*:like:count"
	likeDao := dao.NewLikeDao(context.Background())
	defer likeDao.Close()
	reply, err := likeDao.Do("KEYS", likeKey)
	if err != nil {
		e.ThrowError(e.ErrorRedis)
	}
	likeKeys := reply.([]string)
	reply, err = likeDao.Do("KEYS", likeCountKey)
	likeCountKeys := reply.([]string)
	if err != nil {
		e.ThrowError(e.ErrorRedis)
	}
	for _, key := range likeKeys {
		reply, err := redis.Bool(likeDao.Do("GET", key))
		if err != nil {
			e.ThrowError(e.ErrorRedis)
		}
		strArr := strings.Split(key, ":")
		if err != nil {
			e.ThrowError(e.ErrorRedis)
		}
		like := &model.Like{
			UserId:  strArr[0],
			NoteId:  strArr[2],
			IsLiked: reply,
		}
		exist, err := likeDao.ExistLike(like.UserId, like.NoteId)
		if err != nil {
			e.ThrowError(e.ErrorDataBase)
		}
		if exist {
			err = likeDao.CreateLike(like)
			if err != nil {
				e.ThrowError(e.ErrorDataBase)
			}
		} else {
			err = likeDao.CreateLike(like)
			if err != nil {
				e.ThrowError(e.ErrorDataBase)
			}
		}
	}
	for _, key := range likeCountKeys {
		reply, err := redis.Int(likeDao.Do("GET", key))
		if err != nil {
			e.ThrowError(e.ErrorRedis)
		}
		strArr := strings.Split(key, ":")
		likeCount := &model.LikeCount{
			UserId: strArr[0],
			Count:  reply,
		}
		exist, err := likeDao.ExistLikeCount(likeCount.UserId)
		if err != nil {
			e.ThrowError(e.ErrorDataBase)
		}
		if exist {
			err = likeDao.UpdateLikeCount(likeCount)
			if err != nil {
				e.ThrowError(e.ErrorDataBase)
			}
		} else {
			err = likeDao.CreateLikeCount(likeCount)
			if err != nil {
				e.ThrowError(e.ErrorDataBase)
			}
		}
	}
	e.ThrowSuccess()
}
