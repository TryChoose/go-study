package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

// watch 哨兵监控
func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, //连接地址
		DialTimeout: 5 * time.Second,            //超时时间
	})
	if err != nil {
		fmt.Printf("connect to etcd failed ,err:%v\n", err)
		return
	}
	fmt.Println("etcd connect success")
	defer client.Close()
	ch := client.Watch(context.Background(), "k")
	for wresp := range ch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%v ,key:%v,value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
