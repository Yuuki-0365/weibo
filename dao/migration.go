package dao

import (
	"fmt"
	"weibo/model"
)

func migration() {
	err := _db.Set("gorm_table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Notice{},
			&model.Follow{},
			&model.Favorite{},
			&model.Like{},
			&model.Note{},
			&model.Comment{},
			&model.CommentLike{},
			&model.LikeCount{})
	if err != nil {
		fmt.Println("err=", err)
		return
	}
}
