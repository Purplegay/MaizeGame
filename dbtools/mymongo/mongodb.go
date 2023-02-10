package mymongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var mymongMgr *MyMongo

var (
	MongoUser = "hly"
	Password  = "114514"
)

type MyMongo struct {
	client *mongo.Client
	db     string
}

func GetInstance() *MyMongo {
	once.Do(func() {
		mymongMgr = new(MyMongo)
		mymongMgr.init()
	})

	return mymongMgr
}

func (this *MyMongo) init() {
	var err error
	mongoOption := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoOption.Auth = &options.Credential{
		Username: MongoUser,
		Password: Password,
	}
	this.db = "Meize"

	// 连接到MongoDB
	this.client, err = mongo.Connect(context.TODO(), mongoOption)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = this.client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (this *MyMongo) GetCollection(collectionName string) *mongo.Collection {
	if this.client == nil {
		return nil
	}
	collection := this.client.Database(this.db).Collection(collectionName)
	return collection
}
