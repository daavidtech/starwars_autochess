package match

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type BattleUnit struct {
	Team          int
	UnitID        string
	UnitType      string
	PlayerID      string
	Rank          int
	MaxHP         int
	HP            int
	MaxMana       int
	Mana          int
	AttackRate    int
	AttackRange   int
	AttackDamage  int
	InstantAttack bool
	MoveSpeed     int
	Dead          bool
	X             float32
	Y             float32
	lastAttacked  time.Time
	nextLoc       *Point

	currAttackTarget *BattleUnit
}

func createBattleUnit(unit *Unit, team int, playerID string) *BattleUnit {
	y := unit.Placement.Y

	if team == 2 {
		y = 100 - y
	}

	unitID := uuid.New().String()

	return &BattleUnit{
		UnitID:      unitID,
		UnitType:    unit.UnitType,
		PlayerID:    playerID,
		MaxHP:       100,
		HP:          100,
		MaxMana:     unit.Mana,
		Mana:        unit.Mana,
		AttackRate:  unit.AttackRate,
		AttackRange: 20,
		MoveSpeed:   20,
		Dead:        false,
		Team:        team,
		X:           unit.Placement.X,
		Y:           y,
	}
}

func (battleUnit *BattleUnit) canAttack(now time.Time) bool {
	elapsed := now.Sub(battleUnit.lastAttacked)

	if elapsed.Milliseconds() < int64(battleUnit.AttackRate) {
		return false
	}

	return true
}

func (unit *BattleUnit) calcNextLoc() *Point {
	xDiff := unit.X - unit.currAttackTarget.X
	xDist := int(math.Abs(float64(xDiff)))

	yDiff := unit.Y - unit.currAttackTarget.Y
	yDist := int(math.Abs(float64(yDiff)))

	if unit.isInsideAttackRange(unit.currAttackTarget.X, unit.currAttackTarget.Y) {
		return nil
	}

	if xDist > yDist {
		return &Point{
			X: unit.currAttackTarget.X,
			Y: unit.Y,
		}
	} else {
		return &Point{
			X: unit.X,
			Y: unit.currAttackTarget.Y,
		}
	}
}

func (unit *BattleUnit) moveTowardsNextLoc(delta float32) {
	if unit.nextLoc == nil {
		return
	}

	xDiff := unit.nextLoc.X - unit.X
	xDist := int(math.Abs(float64(xDiff)))

	yDiff := unit.nextLoc.Y - unit.Y
	yDist := int(math.Abs(float64(yDiff)))

	stepSize := float32(unit.MoveSpeed) * delta

	if xDist > yDist {
		stepSize = float32(math.Min(float64(xDist), float64(stepSize)))

		if xDiff < 0 {
			stepSize *= -1
		}

		unit.X += stepSize
	} else {
		stepSize = float32(math.Min(float64(yDist), float64(stepSize)))

		if yDiff < 0 {
			stepSize *= -1
		}

		unit.Y += stepSize
	}
}

func (unit *BattleUnit) isInsideAttackRange(x float32, y float32) bool {
	xdist := math.Abs(float64(unit.X - x))
	ydist := math.Abs(float64(unit.Y - y))

	return xdist <= float64(unit.AttackRange) && ydist <= float64(unit.AttackRange)
}
