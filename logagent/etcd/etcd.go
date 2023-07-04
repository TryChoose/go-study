package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

type LogETCDConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd
var (
	cli *clientv3.Client
)

// 初始化Init
func Init(address []string, time time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time,
	})
	if err != nil {
		fmt.Printf("cli connect failed ,err:%v\n", err)
		return
	}
	return
}

// 从etcd中根据key获取配置项
func GetConf(key string) (LEC []*LogETCDConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed ,err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &LEC)
		if err != nil {
			fmt.Printf("json unmarshal failed ,err:%v\n", err)
			return
		}
	}
	return
}
