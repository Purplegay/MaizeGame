package playermanager

import (
	"rpctest/dep/ticker"
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
	return true
}
