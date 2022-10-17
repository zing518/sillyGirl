package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/cdle/sillyGirl/develop/core"
	"github.com/cdle/sillyGirl/utils"
)

var sillyGirl = core.MakeBucket("sillyGirl")

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc
	core.Init()
	if sillyGirl.GetBool("anti_kasi") {
		go utils.MonitorGoroutine()
	}
	port := sillyGirl.GetString("port", "8080")
	if core.HttpPort != "" {
		port = core.HttpPort
	}
	logs.Info("Http服务已运行(%s)。", sillyGirl.GetString("port", "8080"))
	go core.Server.Run("0.0.0.0:" + port)
	logs.Info("命令怎么失效了？关注频道 https://t.me/kczz2021 获取最新消息，发言关注微信公众号送喜官获取暗号。")
	select {}
}
