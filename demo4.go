package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-02 10:47
* @Description:
**/

type TimePoint struct {
	StartTime int64 `json:"start_time" bson:"start_time"`
	EndTime   int64 `json:"end_time" bson:"end_time"`
}

type LogRecord struct {
	JobName   string    `json:"job_name" bson:"job_name"`
	Command   string    `json:"command" bson:"command"`
	Err       string    `json:"err" bson:"err"`
	Content   string    `json:"content" bson:"content"`
	TimePoint TimePoint `json:"time_point" bson:"time_point"`
}

type FindByJobName struct {
	JobName string `json:"job_name" bson:"job_name"`
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		cond       *FindByJobName
		cursor     *mongo.Cursor
		record     *LogRecord
	)
	// 1, 建立连接
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	opt := options.Client().ApplyURI("mongodb://root:password@localhost:27017")
	if client, err = mongo.Connect(ctx, opt); err != nil {
		panic(err)
	} else {
		ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
		err := client.Ping(ctx2, readpref.Primary())
		if err != nil {
			panic(err)
		}
		database = client.Database("cron")
	}

	// 3, 选择表my_collection
	collection = database.Collection("log")

	cond = &FindByJobName{JobName: "job12"}
	skip := int64(0)
	limit := int64(3)
	if cursor, err = collection.Find(context.TODO(), cond, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}); err != nil {
		fmt.Println(err)
		return
	}

	defer cursor.Close(context.TODO())

	// 6, 遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义一个日志对象
		record = &LogRecord{}

		// 反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		// 把日志行打印出来
		fmt.Println(*record)
	}
}
