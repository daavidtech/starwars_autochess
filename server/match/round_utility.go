package match

// type HandleUnitsResult struct {
// 	alldead bool
// 	events  []MatchEvent
// }

// func doRoundActionsOnUnits(
// 	yourUnits map[string]*BattleUnit,
// 	enemyUnits map[string]*BattleUnit,
// 	delta float32,
// ) HandleUnitsResult {
// 	result := HandleUnitsResult{
// 		alldead: true,
// 	}

// 	now := time.Now()

// 	for _, unit := range yourUnits {
// 		if unit.dead {
// 			continue
// 		}

// 		result.alldead = false

// 		if unit.currAttackTarget == nil {
// 			closestUnit := findClosestUnit(unit, enemyUnits)

// 			unit.currAttackTarget = closestUnit
// 		}

// 		target := unit.currAttackTarget

// 		if target == nil {
// 			continue
// 		}

// 		if unit.isInsideAttackRange(
// 			target.x,
// 			target.y,
// 		) {
// 			if unit.nextLoc != nil {
// 				result.events = append(result.events,
// 					MatchEvent{
// 						UnitArrivedTo: &UnitArrivedTo{
// 							UnitID: unit.unitId,
// 							X:      unit.x,
// 							Y:      unit.y,
// 						},
// 					},
// 				)

// 				unit.nextLoc = nil
// 			}

// 			if !unit.canAttack(now) {
// 				continue
// 			}

// 			unit.lastAttacked = now
// 			target.hp -= 10

// 			if target.hp < 1 {
// 				target.dead = true

// 				closestUnit := findClosestUnit(unit, enemyUnits)

// 				unit.currAttackTarget = closestUnit

// 				result.events = append(result.events, MatchEvent{
// 					UnitDied: &UnitDied{
// 						UnitID: target.unitId,
// 					},
// 				})
// 			}

// 			continue
// 		}

// 		nextLock := unit.calcNextLoc()

// 		if nextLock == nil {
// 			result.events = append(result.events, MatchEvent{
// 				UnitArrivedTo: &UnitArrivedTo{
// 					UnitID: unit.unitId,
// 					X:      unit.x,
// 					Y:      unit.y,
// 				},
// 			})

// 			unit.nextLoc = nil

// 			continue
// 		}

// 		if unit.nextLoc == nil || !nextLock.isEqual(unit.nextLoc) {
// 			unit.nextLoc = nextLock

// 			result.events = append(result.events, MatchEvent{
// 				UnitStartedMovingTo: &UnitStartedMovingTo{
// 					UnitID: unit.unitId,
// 					X:      nextLock.x,
// 					Y:      nextLock.y,
// 				},
// 			})
// 		}

// 		unit.moveTowardsNextLoc(delta)
// 	}

// 	return result
// }

// func findClosestUnit(
// 	yourUnit *BattleUnit,
// 	units map[string]*BattleUnit,
// ) *BattleUnit {
// 	var closest *BattleUnit

// 	x := yourUnit.x
// 	y := yourUnit.y

// 	for _, unit := range units {
// 		if unit.dead {
// 			continue
// 		}

// 		if unit.unitId == yourUnit.unitId {
// 			continue
// 		}

// 		if closest == nil {
// 			closest = unit

// 			continue
// 		}

// 		unit1Distance := calcDist(x, y, closest.x, closest.y)
// 		unit2Distance := calcDist(x, y, unit.x, unit.y)

// 		if unit2Distance >= unit1Distance {
// 			continue
// 		}

// 		closest = unit
// 	}

// 	return closest
// }
