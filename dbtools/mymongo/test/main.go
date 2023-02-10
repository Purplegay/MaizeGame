package main

import (
	"context"
	"fmt"
	"rpctest/dbtools/mymongo"
	playerinfo "rpctest/module/player/entity"
)

func main() {
	player := playerinfo.NewPlayer(1)
	mymongo.GetInstance().SavePlayer(context.TODO(), "Player", player.Uid, player)
	player2 := mymongo.GetInstance().GetPlayer(context.TODO(), "Player", 1)
	fmt.Println(player2)
}
