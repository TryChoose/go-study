package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/*/
redis 连接
*/
var ctx = context.Background()
var rdb *redis.Client

//普通连接模式
func initRedis1() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功！")
}

//操作string类型
func setString() {
	//命令set
	result, err := rdb.Set(ctx, "k1", "20", 0).Result()
	//err = rdb.Set(ctx, "k1", "20", 0).Err()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(result)
	s, err := rdb.Get(ctx, "k1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	////命令setnx
	b, err := rdb.SetNX(ctx, "k2", "20", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(b) //false 已经存在
	//命令setex
	err = rdb.SetEx(ctx, "k3", "30", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
	//setrange 改变从索引开始后面的值
	i, err := rdb.SetRange(ctx, "k2", 1, "11").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	//mset 设置多个值
	s2, err := rdb.MSet(ctx, "k4", "40", "k5", "50").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(s2)
	//mget 取多个值
	strSli, err := rdb.MGet(ctx, "k1", "k2", "k4").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(strSli)
	//append 追加
	k1len, err := rdb.Append(ctx, "k1", "0292").Result()
	fmt.Println(k1len)
	//返回所有的key
	sl, err := rdb.Keys(ctx, "*").Result()
	fmt.Println(sl)
	//遍历string类型key的值
	for _, v := range sl {
		s3, _ := rdb.GetRange(ctx, v, 0, -1).Result()
		fmt.Println(v, s3)
	}
}

//hash 类型的操作命令
func hashOp() {
	//key其实是俗称变量名,这波误以为是key
	i, err := rdb.HSet(ctx, "Mset", map[string]interface{}{"key1": "value1", "key2": "value2"}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	result, err := rdb.HGet(ctx, "Mset", "key1").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	//获取hash的长度
	n, err := rdb.HLen(ctx, "Mset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	//获取全部的hash中的filed以及value
	m, err := rdb.HGetAll(ctx, "Mset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}

/*
相当于是一个栈或者是队列
*/
func listOp01() {
	//从头添加元素
	i, err := rdb.LPush(ctx, "l1", 1, 2, 3, 4, 5, 6).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	strings, err := rdb.LRange(ctx, "l1", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings)
	//设置值
	result, err := rdb.LSet(ctx, "l1", 0, 2).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
func listOp02() {
	//从尾部添加元素
	n, err := rdb.RPush(ctx, "l2", 1, 2, 3, 4, 5).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
func main() {
	initRedis1()
	//setString()
	//hashOp()
	//listOp01()
	listOp02()

}
