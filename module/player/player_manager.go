package playermanager

import (
	"context"
	"rpctest/dbtools/mymongo"
	"rpctest/dep/ticker"
	playerinfo "rpctest/module/player/entity"
	"sync"
	"time"
)

var once sync.Once
var mgr *PlayerManager

type PlayerManager struct {
	activePlayers sync.Map
}

func GetInstance() *PlayerManager {
	once.Do(func() {
		mgr = new(PlayerManager)
		mgr.init()
	})

	return mgr
}

func (this *PlayerManager) init() bool {
	//检查活跃玩家，不活跃则删除缓存
	ticker.GetInstance().AddTicker("TickCheckActivePlayer", time.Hour, func() {
		this.tickCheckActivePlayer()
	})

	//定时保存玩家数据
	ticker.GetInstance().AddTicker("TickSavePlayer", 5*time.Minute, func() {
		this.tickSavePlayer()
	})
	return true
}

func (this *PlayerManager) GetPlayer(uid uint64) *playerinfo.Player {
	if uid == 0 {
		return nil
	}
	v, ok := this.activePlayers.Load(uid)
	if ok {
		return v.(*playerinfo.Player)
	}
	p := mymongo.GetInstance().GetPlayer(context.TODO(), uid)
	if p == nil {
		p = this.createPlayer(uid)
		this.activePlayers.Store(uid, p)
	}
	return p
}

//自动创号
func (this *PlayerManager) createPlayer(uid uint64) *playerinfo.Player {
	if uid == 0 {
		return nil
	}
	p := playerinfo.NewPlayer(uid)
	p.OnCreate()
	mymongo.GetInstance().SavePlayer(context.TODO(), uid, p)
	return p
}
