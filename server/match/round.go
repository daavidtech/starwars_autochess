package match

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

type RoundResult struct {
	whoWon      int
	interrupted bool
}

type Round struct {
	id        string
	player1ID string
	player2ID string
	units     []*BattleUnit

	eventBroker *MatchEventBroker

	ctx context.Context
}

func CreateRound(ctx context.Context, matchEventBroker *MatchEventBroker, units []*BattleUnit) *Round {
	round := &Round{
		id:          uuid.New().String(),
		units:       units,
		eventBroker: matchEventBroker,
		ctx:         ctx,
	}

	return round
}

func (round *Round) findClosestEnemy(yourUnit *BattleUnit) *BattleUnit {
	var closest *BattleUnit

	x := yourUnit.X
	y := yourUnit.Y

	for _, unit := range round.units {
		if unit.Dead {
			continue
		}

		if unit.UnitID == yourUnit.UnitID {
			continue
		}

		if unit.Team == yourUnit.Team {
			continue
		}

		if closest == nil {
			closest = unit

			continue
		}

		unit1Distance := calcDist(x, y, closest.X, closest.Y)
		unit2Distance := calcDist(x, y, unit.X, unit.Y)

		if unit2Distance >= unit1Distance {
			continue
		}

		closest = unit
	}

	return closest
}

type RoundWorkResult struct {
	whoWon int
	events []MatchEvent
}

func (round *Round) work(delta float32) RoundWorkResult {
	result := RoundWorkResult{}

	now := time.Now()

	team1AllDead := true
	team2AllDead := true

	for _, unit := range round.units {
		if unit.Dead {
			continue
		}

		if unit.Team == 1 {
			team1AllDead = false
		}

		if unit.Team == 2 {
			team2AllDead = false
		}

		if unit.currAttackTarget == nil {
			closestUnit := round.findClosestEnemy(unit)

			unit.currAttackTarget = closestUnit
		}

		target := unit.currAttackTarget

		if target == nil {
			continue
		}

		// log.Printf("Unit %v %v target %v %v", unit.X, unit.Y, target.X, target.Y)

		if unit.isInsideAttackRange(
			target.X,
			target.Y,
		) {
			if unit.nextLoc != nil {
				log.Printf("1. Unit arrived to %v %v", unit.X, unit.Y)

				result.events = append(result.events, MatchEvent{
					UnitArrivedTo: &UnitArrivedTo{
						PlayerID: round.player1ID,
						UnitID:   unit.UnitID,
						X:        int(unit.X),
						Y:        int(unit.Y),
					},
				})

				if round.player1ID != round.player2ID {
					invertedY := int(invertY(unit.Y))

					result.events = append(result.events, MatchEvent{
						UnitArrivedTo: &UnitArrivedTo{
							PlayerID: round.player2ID,
							UnitID:   unit.UnitID,
							X:        int(unit.X),
							Y:        invertedY,
						},
					})
				}

				unit.nextLoc = nil
			}

			if !unit.canAttack(now) {
				continue
			}

			log.Printf("Unit is attacking %v", unit.UnitID)

			unit.lastAttacked = now
			target.HP -= unit.AttackDamage

			result.events = append(result.events, MatchEvent{
				BattleUnitHealthChanged: &BattleUnitHealthChanged{
					PlayerID:  round.player1ID,
					UnitID:    target.UnitID,
					NewHealth: target.HP,
				},
			})

			if round.player1ID != round.player2ID {
				result.events = append(result.events, MatchEvent{
					BattleUnitHealthChanged: &BattleUnitHealthChanged{
						PlayerID:  round.player2ID,
						UnitID:    target.UnitID,
						NewHealth: target.HP,
					},
				})
			}

			if target.HP < 1 {
				target.Dead = true

				closestUnit := round.findClosestEnemy(unit)

				unit.currAttackTarget = closestUnit

				log.Printf("Unit died %v", unit.UnitID)

				result.events = append(result.events, MatchEvent{
					UnitDied: &UnitDied{
						PlayerID: round.player1ID,
						UnitID:   target.UnitID,
					},
				})

				if round.player1ID != round.player2ID {
					result.events = append(result.events, MatchEvent{
						UnitDied: &UnitDied{
							PlayerID: round.player2ID,
							UnitID:   target.UnitID,
						},
					})
				}
			}

			continue
		}

		nextLock := unit.calcNextLoc()

		if nextLock == nil {
			log.Printf("2. Unit arrived to %v %v", unit.X, unit.Y)

			result.events = append(result.events, MatchEvent{
				UnitArrivedTo: &UnitArrivedTo{
					PlayerID: round.player1ID,
					UnitID:   unit.UnitID,
					X:        int(unit.X),
					Y:        int(unit.Y),
				},
			})

			if round.player1ID != round.player2ID {
				result.events = append(result.events, MatchEvent{
					UnitArrivedTo: &UnitArrivedTo{
						PlayerID: round.player2ID,
						UnitID:   unit.UnitID,
						X:        int(unit.X),
						Y:        int(invertY(unit.Y)),
					},
				})
			}

			unit.nextLoc = nil

			continue
		}

		if unit.nextLoc == nil || !nextLock.isEqual(unit.nextLoc) {
			unit.nextLoc = nextLock

			log.Printf("Unit started moving to %v %v", nextLock.X, nextLock.Y)

			result.events = append(result.events, MatchEvent{
				UnitStartedMovingTo: &UnitStartedMovingTo{
					PlayerID: round.player1ID,
					UnitID:   unit.UnitID,
					X:        int(nextLock.X),
					Y:        int(nextLock.Y),
				},
			})

			if round.player1ID != round.player2ID {
				result.events = append(result.events, MatchEvent{
					UnitStartedMovingTo: &UnitStartedMovingTo{
						PlayerID: round.player2ID,
						UnitID:   unit.UnitID,
						X:        int(nextLock.X),
						Y:        int(invertY(nextLock.Y)),
					},
				})
			}
		}

		unit.moveTowardsNextLoc(delta)
	}

	if team1AllDead {
		result.whoWon = 2
	}

	if team2AllDead {
		result.whoWon = 1
	}

	return result
}

func (round *Round) run() RoundResult {
	now := time.Now()
	var result RoundWorkResult

	for result.whoWon == 0 {
		elapsed := time.Since(now)
		now = time.Now()

		result = round.work(float32(elapsed.Seconds()))

		if len(result.events) > 0 {
			round.eventBroker.publishEvent(result.events...)
		}

		select {
		case <-time.NewTimer(20 * time.Millisecond).C:
		case <-round.ctx.Done():
			return RoundResult{
				interrupted: true,
			}
		}
	}

	// log.Printf("Team %v won", roundResult.whoWon)

	return RoundResult{
		whoWon: result.whoWon,
	}
}
