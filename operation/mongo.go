package operation

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-16 06:11
* @Description: mongoDB connect pool
**/

const NUM = 20
const TIMEOUT = 2 * time.Second
const URI = "mongodb://localhost:12071"

var client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()
	o := options.Client().ApplyURI(URI)
	o.SetMaxPoolSize(NUM)
	var err error
	client, err = mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetConn(){

}