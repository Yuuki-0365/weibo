package dao

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"weibo/cache"
	"weibo/model"
)

type LikeDao struct {
	redis.Conn
	*gorm.DB
}

func NewLikeDao(ctx context.Context) *LikeDao {
	return &LikeDao{cache.RedisPool.Get(), NewDBClient(ctx)}
}

func (dao *LikeDao) CreateLike(like *model.Like) (err error) {
	err = dao.DB.Model(&model.Like{}).Create(like).Error
	return err
}

func (dao *LikeDao) CreateLikeCount(likeCount *model.LikeCount) (err error) {
	err = dao.DB.Model(&model.LikeCount{}).Create(likeCount).Error
	return err
}

func (dao *LikeDao) ExistLike(userId, noteId string) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Like{}).
		Where("user_id = ? and note_id = ?", userId, noteId).
		Count(&count).Error
	if err != nil {
		return true, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (dao *LikeDao) ExistLikeCount(userId string) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.LikeCount{}).
		Where("user_id = ?", userId).
		Count(&count).Error
	if err != nil {
		return true, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (dao *LikeDao) UpdateLike(like *model.Like) (err error) {
	err = dao.DB.Model(&model.Like{}).Where("user_id = ? and note_id = ?", like.UserId, like.NoteId).
		Update("is_liked = ?", like.IsLiked).Error
	return err
}

func (dao *LikeDao) UpdateLikeCount(likeCount *model.LikeCount) (err error) {
	err = dao.DB.Model(&model.LikeCount{}).Where("user_id = ?", likeCount.UserId).
		Update("count = ?", likeCount.Count).Error
	return err
}
