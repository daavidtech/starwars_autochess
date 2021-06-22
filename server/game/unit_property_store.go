package game

type UnitPropertyStore struct {
	units map[string]UnitProperties
}

func (unitRegister *UnitPropertyStore) SaveUnit(unit UnitProperties) {
	unitRegister.units[unit.UnitType] = unit
}

func (unitRegister *UnitPropertyStore) GetUnit(unitType string) UnitProperties {
	return unitRegister.units[unitType]
}

func (u *UnitPropertyStore) ChooseRandomUnitFromTier(tier int) UnitProperties {
	return UnitProperties{}
}
