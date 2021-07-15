package match

import (
	"math/rand"
)

type ShopUnit struct {
	ID int
	UnitProperties
}

type Shop struct {
	size  int
	units map[int]ShopUnit

	TierProbabilities *TierProbabilities
	UnitPropertyStore *UnitPropertyStore
}

func (shop *Shop) SetSize(s int) {
	shop.size = s

}

func NewShop() *Shop {
	return &Shop{
		units: make(map[int]ShopUnit),
	}
}

func (shop *Shop) Fill(level int) ShopRefilled {
	shop.units = make(map[int]ShopUnit)

	shopRefilled := ShopRefilled{
		ShopUnits: []ShopUnit{},
	}

	if shop.TierProbabilities == nil {
		return shopRefilled
	}

	if shop.UnitPropertyStore == nil {
		return shopRefilled
	}

	if shop.UnitPropertyStore.CountUnits() == 0 {
		return shopRefilled
	}

	for i := 0; i < shop.size; i++ {
		id := i + 1

		probabilities := shop.TierProbabilities.PickLevel(level)

		tier := chooseRandomTier(probabilities)

		unitProps := shop.UnitPropertyStore.ChooseRandomUnitFromTier(tier)

		shopUnit := ShopUnit{
			ID:             id,
			UnitProperties: unitProps,
		}

		shop.units[id] = shopUnit

		shopRefilled.ShopUnits = append(shopRefilled.ShopUnits, shopUnit)
	}

	return shopRefilled
}

func (shop *Shop) GetUnits() []ShopUnit {
	shopUnits := []ShopUnit{}

	for _, unit := range shop.units {
		shopUnits = append(shopUnits, unit)
	}

	return shopUnits
}

func (shop *Shop) GetUnitsCount() int {
	return len(shop.units)
}

func (shop *Shop) Pick(id int) ShopUnit {
	unit := shop.units[id]

	delete(shop.units, id)

	return unit
}

func chooseRandomTier(prob []int) int {
	l := len(prob)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return 1
	}

	i := rand.Intn(100)

	return chooseTier(prob, i)
}

func chooseTier(prob []int, r int) int {
	start := 0

	for index, p := range prob {
		if p == 0 {
			continue
		}

		if r < start || r > start+p {
			start += p

			continue
		}

		return index + 1
	}

	return 0
}
