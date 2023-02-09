package entity

import "time"

type Player struct {
	Uid      uint64
	UpdateAt int64
	dirty    bool
}

func NewPlayer(uid uint64) *Player {
	p := new(Player)
	p.Uid = uid
	return p
}

func (this *Player) Update() {
	this.UpdateAt = time.Now().Unix()
	this.dirty = true
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
