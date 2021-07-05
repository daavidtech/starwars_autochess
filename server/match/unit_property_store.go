package match

import "math/rand"

type UnitPropertyStore struct {
	units          map[string]UnitProperties
	unitTierCounts map[int]int
}

func NewUnitPropertyStore() UnitPropertyStore {
	return UnitPropertyStore{
		units:          make(map[string]UnitProperties),
		unitTierCounts: make(map[int]int),
	}
}

func (unitStore *UnitPropertyStore) CountUnits() int {
	return len(unitStore.units)
}

func (unitRegister *UnitPropertyStore) SaveUnit(unit UnitProperties) {
	unitRegister.units[unit.UnitType] = unit

	count := unitRegister.unitTierCounts[unit.Tier]

	count += 1

	unitRegister.unitTierCounts[unit.Tier] = count

}

func (unitRegister *UnitPropertyStore) GetUnit(unitType string) UnitProperties {
	return unitRegister.units[unitType]
}

func (u *UnitPropertyStore) ChooseRandomUnitFromTier(tier int) UnitProperties {
	count := u.unitTierCounts[tier]

	if count == 0 {
		return UnitProperties{}
	}

	random := rand.Intn(count)

	i := 0

	for _, u := range u.units {
		if i != random {
			i += 1

			continue
		}

		return u
	}

	return UnitProperties{}
}
