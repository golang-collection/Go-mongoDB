package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-16 06:00
* @Description:
**/

var client *mongo.Client

func main() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("success")
	}

}
