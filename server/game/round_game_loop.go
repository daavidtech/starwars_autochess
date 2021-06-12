package game

import "math"

type RoundUnit struct {
	id           string
	unitId       string
	rank         int
	maxHP        int
	maxMana      int
	hp           int
	mana         int
	xCoor        int
	yCoor        int
	attackRange  int
	attackSpeed  int
	attackDamage int

	instantAttack bool

	sleep int

	dead bool

	sinceLastAttack int

	currentAttackTarget *RoundUnit
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
		roundGameLoop.work(1)
	}
}

func (roundGameLoop *RoundGameLoop) work(delta float32) {
	findAttackTargetForUnits(roundGameLoop.playerOneUnits, roundGameLoop.playerTwoUnits)
	findAttackTargetForUnits(roundGameLoop.playerTwoUnits, roundGameLoop.playerOneUnits)

	handleIsInsideAttackRange(roundGameLoop.playerOneUnits)
	handleIsInsideAttackRange(roundGameLoop.playerTwoUnits)

	moveUnits(roundGameLoop.playerOneUnits, delta)
	moveUnits(roundGameLoop.playerTwoUnits, delta)

	// for _, unit := range roundGameLoop.playerTwoUnits {

	// }
}

func handleIsInsideAttackRange(units map[string]*RoundUnit) {
	for _, unit := range units {
		target := unit.currentAttackTarget

		if !isInsideAttackRange(
			unit.xCoor,
			unit.yCoor,
			target.xCoor,
			target.yCoor,
			unit.attackRange,
		) {
			continue
		}

		target.hp -= unit.attackDamage

		if target.hp > 0 {
			continue
		}

		unit.dead = true
		delete(units, unit.id)
	}
}

func moveUnits(units map[string]*RoundUnit, delta float32) {
	for _, unit := range units {
		xDiff := unit.xCoor - unit.currentAttackTarget.xCoor
		xDist := int(math.Abs(float64(xDiff)))

		yDiff := unit.yCoor - unit.currentAttackTarget.yCoor
		yDist := int(math.Abs(float64(yDiff)))

		if xDist > yDist {
			unit.yCoor += int(float32(yDist) * delta)
		} else {
			unit.xCoor += int(float32(xDist) * delta)
		}
	}
}

func isInsideAttackRange(yourX int, yourY int, x int, y int, attackRange int) bool {
	if math.Abs(float64(yourX-x)) <= float64(attackRange) {
		return true
	}

	if math.Abs(float64(yourY-y)) <= float64(attackRange) {
		return true
	}

	return false
}

func findAttackTargetForUnits(units map[string]*RoundUnit, enemyUnits map[string]*RoundUnit) {
	for _, unit := range units {
		if unit.currentAttackTarget == nil {
			unit := findClosestUnit(unit, enemyUnits)

			unit.currentAttackTarget = unit
		}
	}
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
