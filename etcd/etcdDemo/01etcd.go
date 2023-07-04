package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

// put和get操作
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
	//put操作
	str := `[{"path":"E:/temp/ng.log","topic":"ng1"},{"path":"E:/temp/redis.log","topic":"re1"}]`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = client.Put(ctx, "k", str)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed err:%v\n", err)
		return
	}
	//get操作
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, "k")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed err:%v\n", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s:%s\n", kv.Key, kv.Value)
	}
}
