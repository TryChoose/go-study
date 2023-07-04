package main

import (
	"csgo/Base/myLogger"
)

func main() {
	log := myLogger.NewFileLogger("error", "./", "01", 10*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条Fatal日志")
		//time.Sleep(2 * time.Second)
	}
}
