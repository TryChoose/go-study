package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoDB *mongo.Database

func initDB(ctx context.Context, dbName string) error {
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//检查是否连接成功
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		return err
	}
	fmt.Println("Connected to MongoDB")
	db := client.Database(dbName)
	mongoDB = db
	return nil
}
func Create(ctx context.Context, collectionName string) (err error) {
	err = mongoDB.CreateCollection(ctx, collectionName)
	if err != nil {
		fmt.Println("Error creating MongoDB", err)
		return
	}
	fmt.Println("Success creating MongoDB")
	return nil
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := initDB(ctx, "test1")
	if err != nil {
		fmt.Println("初始化数据库失败!")
		return
	}
	err = Create(ctx, "test1")
	if err != nil {
		fmt.Println("创建集合失败!", err)
		return
	}
	//collections, err := db.ListCollectionNames(ctx, bson.D{})
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//for _, coll := range collections {
	//	fmt.Println(coll)
	//}
	//fmt.Println(collections)
	//cur, err := db.Collection("unicorns").Find(ctx, bson.D{})
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//m1 := make([]map[string]interface{}, 0)
	//for cur.Next(ctx) {
	//	m := make(map[string]interface{})
	//	err = cur.Decode(&m)
	//	if err != nil {
	//		fmt.Println("error", err)
	//	}
	//	m1 = append(m1, m)
	//}
	//for _, m := range m1 {
	//	fmt.Println(m)
	//}
	////fmt.Println(m1)
	//err = db.Drop(ctx)
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//fmt.Println("删除成功")
}
