package schema

type UnitList struct {
	Units map[int32]*UnitCard

	baseSchema
}

func InitUnitList() *UnitList {
	list := &UnitList{
		Units: make(map[int32]*UnitCard),
	}
	return list
}

type UnitCard struct {
	CardId int32 `bson:"CardId,omitempty" msgpack:"CardId,omitempty"`
	Star   int32 `bson:"Star,omitempty" msgpack:"Star,omitempty"`
	Level  int32 `bson:"Level,omitempty" msgpack:"Level,omitempty"`
}

func (this *UnitList) AddCards(unit *UnitCard) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.Units[unit.CardId]; ok {
		return false
	}

	this.Units[unit.CardId] = unit
	return true
}
