package match

import "testing"

func TestChoosesRandomUnit(t *testing.T) {
	unitStore := NewUnitPropertyStore()

	unitStore.SaveUnit(UnitProperties{
		UnitType: "unit_droid",
		Tier:     1,
		HP:       100,
	})

	unitProperties := unitStore.ChooseRandomUnitFromTier(1)

	if unitProperties.Tier != 1 {
		t.Error("Wrong tier")
	}
}
