package match

import "testing"

func TestWhenProbArrayIsEmpty(t *testing.T) {
	index := chooseRandomTier([]int{})

	if index != 0 {
		t.Errorf("Index is not zero")
	}
}

func TestWhenProbArrayLenthIsOne(t *testing.T) {
	index := chooseRandomTier([]int{1})

	if index != 1 {
		t.Errorf("Index is not one")
	}
}

func TestChoosesFirstTier(t *testing.T) {
	index := chooseTier([]int{30, 30, 30}, 23)

	if index != 1 {
		t.Errorf("Index is not one")
	}
}

func TestChoosesSecondTier(t *testing.T) {
	index := chooseTier([]int{30, 30, 30}, 34)

	if index != 2 {
		t.Errorf("Index is not two")
	}
}

func TestChoosesThirdTier(t *testing.T) {
	index := chooseTier([]int{30, 30, 30}, 67)

	if index != 3 {
		t.Errorf("Index is not three")
	}
}

func TestPicksOneUnit(t *testing.T) {
	shop := Shop{
		units: []ShopUnit{ShopUnit{UnitType: "unit_clone"}, ShopUnit{UnitType: "unit_droid"}},
	}

	if shop.GetUnitsCount() != 2 {
		t.Errorf("Shop units count is invalid")
	}

	units := shop.GetUnits()

	if len(units) != 2 {
		t.Errorf("Shop units is invalid")
	}

	unit := shop.Pick(0)

	if unit.UnitType != "UNIT1" {
		t.Errorf("Pick gives wrong unit")
	}

}
