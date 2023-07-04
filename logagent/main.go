package main

import (
	"fmt"
	"go-study/logagent/config"
	"go-study/logagent/etcd"
	"go-study/logagent/kafka"
	"gopkg.in/ini.v1"
	"time"
)

var (
	conf = new(config.AppConf)
)

//func run() {
//	//1.读取文件
//	ch := tail.ReadChannel()
//	for {
//		select {
//		case line := <-ch:
//			kafka.SendToKafka(conf.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//	//2.发送到kafka
//}

// logAgent 入口
func main() {
	//0.加载配置文件
	err := ini.MapTo(conf, "./config/my.ini")
	if err != nil {
		fmt.Println("加载进入结构体失败!")
		return
	}
	//1.初始化kafka连接
	err = kafka.Init([]string{conf.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success!")
	//2.初始化etcd
	err = etcd.Init([]string{conf.EtcdConf.Address}, time.Duration(conf.EtcdConf.Time)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed ,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success!")

	//2.1获取配置信息
	lec, err := etcd.GetConf("k")
	if err != nil {
		fmt.Printf("get conf failed,err:%v\n", err)
		return
	}
	for _, v := range lec {
		fmt.Printf("path:%v topic:%v\n", v.Path, v.Topic)
	}
	////2.打开日志文件收集日志
	//err = tail.Init(conf.TailLogConf.Filename)
	//if err != nil {
	//	fmt.Printf("init tail failed,err", err)
	//	return
	//}
	//run()

}
