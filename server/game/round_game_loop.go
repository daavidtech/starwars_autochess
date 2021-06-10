package game

type RoundUnit struct {
	id      string
	unitId  string
	rank    int
	maxHP   int
	maxMana int
	hp      int
	mana    int
	xCoor   int
	yCoor   int

	sinceLastAttack int

	currentAttackTarget string
}

type RoundGameLoop struct {
	playerOneUnits map[string]*RoundUnit
	playerTwoUnits map[string]*RoundUnit
}

func (roundGameLoop *RoundGameLoop) OnFinish() <-chan interface{} {
	return make(<-chan interface{}, 1)
}

func (roundGameLoop *RoundGameLoop) Run() {
	for {
		roundGameLoop.work()
	}
}

func (roundGameLoop *RoundGameLoop) work() {
	for _, unit := range roundGameLoop.playerOneUnits {
		if unit.currentAttackTarget == "" {
			unit := findClosestUnit(unit, roundGameLoop.playerTwoUnits)

			unit.currentAttackTarget = unit.id
		}
	}

	// for _, unit := range roundGameLoop.playerTwoUnits {

	// }
}

func findClosestUnit(
	yourUnit *RoundUnit,
	units map[string]*RoundUnit,
) *RoundUnit {
	var closest *RoundUnit

	x := yourUnit.xCoor
	y := yourUnit.yCoor

	for _, unit := range units {
		if unit.id == yourUnit.id {
			continue
		}

		if closest == nil {
			closest = unit

			continue
		}

		unit1Distance := calcDist(x, y, closest.xCoor, closest.yCoor)
		unit2Distance := calcDist(x, y, unit.xCoor, unit.yCoor)

		if unit2Distance >= unit1Distance {
			continue
		}

		closest = unit
	}

	return closest
}
