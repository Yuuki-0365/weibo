package main

import (
	"github.com/robfig/cron/v3"
	"weibo/conf"
	"weibo/router"
	"weibo/service"
)

func main() {
	c := cron.New()
	c.AddFunc("30 * * * *", func() {
		var likeService service.LikeService
		likeService.ReadDataFromRedis()
	})
	c.Start()
	// 初始化配置
	conf.Init()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
