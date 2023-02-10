package mymongo

import (
	"context"
	playerinfo "rpctest/module/player/entity"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//获取玩家
func (this *MyMongo) GetPlayer(ctx context.Context, colName string, roleId uint64) *playerinfo.Player {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := this.GetCollection(colName)
	if collection == nil {
		log.Println("GetCollection Nil")
	}

	ret, err := collection.Find(ctx, bson.D{
		{Key: "Uid", Value: roleId},
	})
	if err != nil {
		log.Println("GetPlayer Error:", err)
		return nil
	}
	if ret == nil {
		log.Println("GetPlayer ret Nil!")
		return nil
	}

	player := playerinfo.NewPlayer(roleId)
	for ret.Next(ctx) {
		if err := bson.Unmarshal(ret.Current, player); err != nil {
			log.Println("bson.Unmarshal Error:", err)
			return nil
		} else {
			return player
		}
	}
	return nil
}

//保存玩家
func (this *MyMongo) SavePlayer(ctx context.Context, colName string, roleId uint64, player interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := this.GetCollection(colName)
	if collection == nil {
		log.Println("GetCollection Nil")
		return false
	}

	upset := true
	_, err := collection.ReplaceOne(ctx, bson.D{{Key: "Uid", Value: roleId}}, player, &options.ReplaceOptions{Upsert: &upset})
	if err != nil {
		log.Println("ReplaceOne Error:", err)
		return false
	}
	return true
}
