package dao

import (
	"context"
	"gorm.io/gorm"
	"weibo/model"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDb(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeById(id int64) (notice *model.Notice, err error) {
	err = dao.DB.Table("notice").
		Where("id=?", id).
		First(&notice).Error
	return notice, err
}
