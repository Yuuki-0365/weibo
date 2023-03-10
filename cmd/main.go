package main

import (
	"SmallRedBook/conf"
	"SmallRedBook/router"
	"SmallRedBook/service"
	"github.com/robfig/cron/v3"
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
