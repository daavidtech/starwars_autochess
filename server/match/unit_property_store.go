package match

import "math/rand"

type UnitPropertyStore struct {
	units          []UnitProperties
	unitTierCounts map[int]int
}

func NewUnitPropertyStore() UnitPropertyStore {
	return UnitPropertyStore{
		units: []UnitProperties{},
	}
}

func (unitStore *UnitPropertyStore) CountUnits() int {
	return len(unitStore.units)
}

func (unitPropStore *UnitPropertyStore) SaveUnit(newUnit UnitProperties) {
	if newUnit.HP == 0 {
		panic("Unit cannot have zero hp")
	}

	if newUnit.Rank < 1 && newUnit.Rank > 3 {
		panic("Invalid rank")
	}

	units := []UnitProperties{}

	for _, unit := range unitPropStore.units {
		if unit.UnitType == newUnit.UnitType && unit.Rank == newUnit.Rank {
			continue
		}

		units = append(units, unit)
	}

	units = append(units, newUnit)

	unitPropStore.units = units
}

func (u *UnitPropertyStore) PickRandom(tier int) (UnitProperties, bool) {
	count := u.CountTierUnits(tier)

	if count == 0 {
		return UnitProperties{}, false
	}

	random := rand.Intn(count)

	i := 0

	for _, u := range u.units {
		if u.Rank != 1 {
			continue
		}

		if i != random {
			i += 1

			continue
		}

		return u, true
	}

	return UnitProperties{}, false
}

func (unitPropStore *UnitPropertyStore) CountTierUnits(tier int) int {
	count := 0

	for _, unit := range unitPropStore.units {
		if unit.Tier != tier || unit.Rank != 1 {
			continue
		}

		count += 1
	}

	return count
}

func (unitPropStore *UnitPropertyStore) FindProps(unitType string, rank int) (UnitProperties, bool) {
	for _, unitProp := range unitPropStore.units {
		if unitProp.UnitType == unitType && unitProp.Rank == rank {
			return unitProp, true
		}
	}

	return UnitProperties{}, false
}
