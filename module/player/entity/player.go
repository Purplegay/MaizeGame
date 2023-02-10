package playerinfo

import (
	"rpctest/schema"
	"time"
)

type Player struct {
	Uid      uint64           `bson:"Uid,omitempty" msgpack:"Uid,omitempty"`
	UpdateAt int64            `bson:"UpdateAt,omitempty" msgpack:"UpdateAt,omitempty"`
	Units    *schema.UnitList `bson:"Units,omitempty" msgpack:"Units,omitempty"`
	dirty    bool
}

func NewPlayer(uid uint64) *Player {
	p := new(Player)
	p.Uid = uid
	p.UpdateAt = time.Now().Unix()

	p.Units = schema.InitUnitList()
	return p
}

//活跃刷新
func (this *Player) ActiveRefresh() {
	this.UpdateAt = time.Now().Unix()
}

func (this *Player) Update() {
	this.UpdateAt = time.Now().Unix()
	this.dirty = true
}
func (this *Player) GetDirty() bool {
	return this.dirty
}

func (this *Player) SetDirty(set bool) {
	this.dirty = set
}

//检查是否需要清理
func (this *Player) CheckClear() bool {
	if this.dirty {
		return false
	}
	//3天没更新则清理
	return time.Now().Unix()-this.UpdateAt > 3600*24*3
}
