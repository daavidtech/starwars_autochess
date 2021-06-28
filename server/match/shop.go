package match

import (
	"math/rand"
)

type ShopUnit struct {
	UnitType   string
	Tier       int
	HP         int
	Mana       int
	AttackRate int
}

type Shop struct {
	size  int
	units []ShopUnit

	tierProbabilities TierProbabilities
	unitPropertyStore UnitPropertyStore
}

func (shop *Shop) SetSize(s int) {
	shop.size = s

}

func CreateShop() Shop {
	return Shop{}
}

func (shop *Shop) Fill(level int) {
	shop.units = []ShopUnit{}

	for i := 0; i < shop.size; i++ {
		probabilities := shop.tierProbabilities.PickLevel(level)

		tier := chooseRandomTier(probabilities)

		unit := shop.unitPropertyStore.ChooseRandomUnitFromTier(tier)

		shop.units = append(shop.units, ShopUnit{
			UnitType: unit.UnitType,
		})
	}
}

func (shop *Shop) GetUnits() []ShopUnit {
	return shop.units
}

func (shop *Shop) GetUnitsCount() int {
	return len(shop.units)
}

func (shop *Shop) Pick(index int) ShopUnit {
	unit := shop.units[index]

	shop.units = append(shop.units[:index], shop.units[index+1:]...)

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
