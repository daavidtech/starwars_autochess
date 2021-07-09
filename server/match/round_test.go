package match

// func Test_pathfinding_test(t *testing.T) {
// 	round := Round{
// 		units: []*BattleUnit{
// 			&BattleUnit{
// 				UnitID:      "1",
// 				X:           10,
// 				Y:           10,
// 				Team:        1,
// 				MoveSpeed:   10,
// 				AttackRange: 10,
// 				HP:          1000,
// 			},
// 			&BattleUnit{
// 				UnitID:      "2",
// 				X:           10,
// 				Y:           80,
// 				Team:        2,
// 				MoveSpeed:   10,
// 				AttackRange: 10,
// 				HP:          1000,
// 			},
// 		},
// 	}

// 	now := time.Now()

// 	for {
// 		elapsed := time.Since(now)
// 		now = time.Now()

// 		result := round.work(float32(elapsed.Seconds()))

// 		if result.whoWon == 1 {
// 			log.Println("Team 1 won")
// 			break
// 		}

// 		if result.whoWon == 2 {
// 			log.Println("Team 2 won")
// 			break
// 		}

// 		<-time.NewTimer(20 * time.Millisecond).C
// 	}
// }
