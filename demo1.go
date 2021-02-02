package main

/**
* @Author: super
* @Date: 2021-02-02 09:07
* @Description: 建立连接
**/

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)
	// 1, 建立连接
	if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"),
		options.Client().SetConnectTimeout(5),
		options.Client().SetAuth(options.Credential{
			Username: "root",
			Password: "password",
		})); err != nil {
		fmt.Println(err)
		return
	}

	// 2, 选择数据库my_db
	database = client.Database("my_db")

	// 3, 选择表my_collection
	collection = database.Collection("my_collection")

	collection = collection
}
