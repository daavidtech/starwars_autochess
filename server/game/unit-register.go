package game

type UnitRegister struct {
	units map[string]Unit
}

func (unitRegister *UnitRegister) SaveUnit(unit Unit) {
	unitRegister.units[unit.unitId] = unit
}

func (unitRegister *UnitRegister) GetUnit(unitId string) Unit {
	return unitRegister.units[unitId]
}

func (u *UnitRegister) ChooseRandomUnitFromTier(tier int) Unit {
	return Unit{}
}
