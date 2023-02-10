package playermanager

import playerinfo "rpctest/module/player/entity"

//检查活跃玩家，不活跃则删除缓存
func (this *PlayerManager) tickCheckActivePlayer() {
	this.activePlayers.Range(func(key, value any) bool {
		player, ok := value.(*playerinfo.Player)
		if !ok {
			return true
		}
		if player.CheckClear() {
			this.activePlayers.Delete(key)
		}
		return true
	})
}
