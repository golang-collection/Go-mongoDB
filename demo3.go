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
* @Date: 2021-02-02 10:40
* @Description: 批量插入数据
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

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		result     *mongo.InsertManyResult
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

	record = &LogRecord{
		JobName:   "job12",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	logArray := []interface{}{record, record, record}

	result, err = collection.InsertMany(context.TODO(), logArray)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.InsertedIDs)

}
