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
* @Date: 2021-02-02 11:20
* @Description: 删除操作
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

// startTime小于某时间
// {"$lt": timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

// {"timePoint.startTime": {"$lt": timestamp} }
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		delCond    *DeleteCond
		delResult  *mongo.DeleteResult
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

	// 4, 要删除开始时间早于当前时间的所有日志($lt是less than)
	//  delete({"timePoint.startTime": {"$lt": 当前时间}})
	delCond = &DeleteCond{beforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	// 执行删除
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数:", delResult.DeletedCount)
}
