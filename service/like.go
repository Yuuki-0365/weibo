package service

import (
	"SmallRedBook/dao"
	"SmallRedBook/tool"
	"SmallRedBook/tool/e"
	"context"
	"fmt"
)

type LikeService struct {
	NoteId string `json:"note_id" form:"note_id"`
}

func (service *LikeService) LikeNote(ctx context.Context, userId string) tool.Response {
	key := userId + ":like:" + service.NoteId
	county := userId + ":like:count"
	likeDao := dao.NewLikeDao(ctx)
	value, err := likeDao.Do("GET", key)
	if err != nil {
		e.ThrowError(e.ErrorRedis)
	}
	if value == nil {
		_, err = likeDao.Do("SET", key, 1)
		if err != nil {
			e.ThrowError(e.ErrorRedis)
		}
		_, err = likeDao.Do("incr", county)
		if err != nil {
			e.ThrowError(e.ErrorRedis)
		}
	} else {
		str := fmt.Sprintf("%s", value)
		if str == "1" {
			_, err = likeDao.Do("SET", key, 0)
			if err != nil {
				e.ThrowError(e.ErrorRedis)
			}
			_, err = likeDao.Do("decr", county)
			if err != nil {
				e.ThrowError(e.ErrorRedis)
			}
		} else {
			_, err = likeDao.Do("SET", key, 1)
			if err != nil {
				e.ThrowError(e.ErrorRedis)
			}
			_, err = likeDao.Do("incr", county)
			if err != nil {
				e.ThrowError(e.ErrorRedis)
			}
		}
	}
	return tool.Response{
		Status: e.Success,
		Msg:    e.GetMsg(e.Success),
	}
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
