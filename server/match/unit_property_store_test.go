package match

import "testing"

func TestChoosesRandomUnit(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitProperties, _ := unitStore.PickRandom(1)

	if unitProperties.Tier != 1 {
		t.Error("Wrong tier")
	}
}

func Test_trying_to_pick_when_no_units_returns_success_false(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	_, success := unitStore.PickRandom(1)

	if success != false {
		t.Error("Success should be false")
	}
}

func Test_chooses_only_rank1(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_clone",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     2,
		Tier:     1,
		HP:       100,
	})

	unitProperties, success := unitStore.PickRandom(1)

	if success == false {
		t.Error("Picking failed")
	}

	if unitProperties.Tier != 1 {
		t.Error("Wrong tier")
	}
}

func Test_count_tier_units(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_clone",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     2,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_darth_vader",
		Rank:     1,
		Tier:     5,
		HP:       100,
	})

	count := unitStore.CountTierUnits(1)

	if count != 2 {
		t.Errorf("Counted invalid amount of units %v", count)
	}
}

func Test_adds_new_unit(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	count := unitStore.CountUnits()

	if count != 1 {
		t.Errorf("Invalid number of units %v", count)
	}
}

func Test_updates_existing_unit(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       200,
	})

	unit := unitStore.units[0]

	if unit.HP != 200 {
		t.Error("Unit is not upgraded")
	}
}

func Test_find_props(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     1,
		Tier:     1,
		HP:       100,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Rank:     2,
		Tier:     1,
		HP:       200,
	})

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_clone",
		Rank:     1,
		Tier:     1,
		HP:       42,
	})

	unitProps, success := unitStore.FindProps("unit_clone", 1)

	if unitProps.Rank != 1 &&
		unitProps.Tier != 1 &&
		unitProps.HP != 42 &&
		unitProps.UnitType != "unit_clone" &&
		success == false {

		t.Error("Unit not found")
	}
}
