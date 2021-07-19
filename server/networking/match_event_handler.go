package networking

import (
	"log"

	"github.com/daavidtech/starwars_autochess/match"
)

func handle_match_events(ws WSWriter, currentMatch *match.Match, playerID string) {
	eventBroker := currentMatch.GetEventBroker()

	ch := eventBroker.Subscribe(currentMatch.GetID())

	defer eventBroker.Unsubscribe(ch)

	for event := range ch {
		// if event.NewBarrackUnit != nil {

		// }

		var err error

		if event.ShopRefilled != nil && event.ShopRefilled.PlayerID == playerID {
			shopRefilled := ShopRefilled{
				ShopUnits: []ShopUnit{},
			}

			for _, shopUnit := range event.ShopRefilled.ShopUnits {
				shopRefilled.ShopUnits = append(shopRefilled.ShopUnits, ShopUnit{
					ID:       shopUnit.ID,
					UnitType: shopUnit.UnitType,
					Level:    shopUnit.Level,
					HP:       shopUnit.HP,
					Mana:     shopUnit.Mana,
					Rank:     shopUnit.Rank,
					Cost:     shopUnit.Cost,
				})
			}

			msg := MessageToClient{
				ShopRefilled: &shopRefilled,
			}

			log.Println("Sending shopRefilled to client")

			ws.WriteJSON(msg)
		}

		if event.CountdownStarted != nil {
			ws.WriteJSON(MessageToClient{
				CountdownStarted: &CountdownStarted{
					StartValue: event.CountdownStarted.StartValue,
					Interval:   event.CountdownStarted.Interval,
				},
			})
		}

		if event.ShopUnitRemoved != nil && event.ShopUnitRemoved.PlayerID == playerID {
			log.Println("Sending shop unit removed")

			ws.WriteJSON(MessageToClient{
				ShopUnitRemoved: &ShopUnitRemoved{
					ShopUnitID: event.ShopUnitRemoved.ShopUnitID,
				},
			})
		}

		if event.BarrackUnitAdded != nil && event.BarrackUnitAdded.PlayerID == playerID {
			if event.BarrackUnitAdded.PlayerID == playerID {
				unitAdded := UnitAdded{
					UnitID: event.BarrackUnitAdded.UnitID,
					UnitProperties: UnitProperties{
						UnitType:   event.BarrackUnitAdded.UnitType,
						Rank:       event.BarrackUnitAdded.Rank,
						HP:         event.BarrackUnitAdded.HP,
						Mana:       event.BarrackUnitAdded.Mana,
						AttackRate: event.BarrackUnitAdded.AttackRate,
					},
				}

				msg := MessageToClient{
					UnitAdded: &unitAdded,
				}

				log.Println("Seding unitAdded to client")

				ws.WriteJSON(msg)
			}
		}

		if event.BarrackUnitRemoved != nil && event.BarrackUnitRemoved.PlayerID == playerID {
			if event.BarrackUnitRemoved.PlayerID == playerID {
				unitRemoved := UnitRemoved{
					UnitID: event.BarrackUnitRemoved.UnitID,
				}

				msg := MessageToClient{
					UnitRemoved: &unitRemoved,
				}

				log.Println("Sending unitRemoved to client")

				ws.WriteJSON(msg)
			}
		}

		if event.BarrackUnitUpgraded != nil && event.BarrackUnitUpgraded.PlayerID == playerID {
			if event.BarrackUnitUpgraded.PlayerID == playerID {
				unitUpgraded := UnitUpgraded{
					UnitID: event.BarrackUnitUpgraded.UnitID,
					UnitProperties: UnitProperties{
						Rank:       event.BarrackUnitUpgraded.Rank,
						HP:         event.BarrackUnitUpgraded.HP,
						Mana:       event.BarrackUnitUpgraded.Mana,
						AttackRate: event.BarrackUnitUpgraded.AttackRate,
						UnitType:   event.BarrackUnitUpgraded.UnitType,
					},
				}

				log.Println("Sending unitUpgraded to client")

				ws.WriteJSON(MessageToClient{
					UnitUpgraded: &unitUpgraded,
				})
			}
		}

		if event.UnitPlaced != nil && event.UnitPlaced.PlayerID == playerID {

			log.Println("Sending unit placed to client")

			unitPlaced := UnitPlaced{
				UnitID: event.UnitPlaced.UnitID,
				X:      event.UnitPlaced.X,
				Y:      event.UnitPlaced.Y,
			}

			err = ws.WriteJSON(MessageToClient{
				UnitPlaced: &unitPlaced,
			})

		}

		if event.PhaseChanged != nil {
			log.Println("Sending phase changed to client")

			if event.PhaseChanged.MatchPhase == match.EndPhase {
				err = ws.WriteJSON(MessageToClient{
					GamePhaseChanged: &GamePhaseChanged{
						GamePhase: MainMenuPhase,
					},
				})

				return
			} else {
				err = ws.WriteJSON(MessageToClient{
					GamePhaseChanged: &GamePhaseChanged{
						GamePhase: convertMatchPhase(event.PhaseChanged.MatchPhase),
					},
				})
			}
		}

		if event.UnitStartedMovingTo != nil && event.UnitStartedMovingTo.PlayerID == playerID {
			log.Printf("Sending unit %v started moving to %v %v",
				event.UnitStartedMovingTo.UnitID,
				event.UnitStartedMovingTo.X,
				event.UnitStartedMovingTo.Y)

			err = ws.WriteJSON(MessageToClient{
				UnitStartedMovingTo: &UnitStartedMovingTo{
					UnitID: event.UnitStartedMovingTo.UnitID,
					X:      event.UnitStartedMovingTo.X,
					Y:      event.UnitStartedMovingTo.Y,
				},
			})

		}

		if event.UnitArrivedTo != nil && event.UnitArrivedTo.PlayerID == playerID {
			log.Printf("Sending unit %v arrived to", event.UnitArrivedTo.UnitID)

			err = ws.WriteJSON(MessageToClient{
				UnitArrivedTo: &UnitArrivedTo{
					UnitID: event.UnitArrivedTo.UnitID,
					X:      event.UnitArrivedTo.X,
					Y:      event.UnitArrivedTo.Y,
				},
			})

		}

		if event.UnitDied != nil && event.UnitDied.PlayerID == playerID {

			log.Printf("Sending unit %v died to client", event.UnitDied.UnitID)

			err = ws.WriteJSON(MessageToClient{
				UnitDied: &UnitDied{
					UnitID: event.UnitDied.UnitID,
				},
			})

		}

		if event.RoundCreated != nil && event.RoundCreated.PlayerID == playerID {

			roundCreated := RoundCreated{
				PlayerID: playerID,
				Units:    []BattleUnit{},
			}

			for _, unit := range event.RoundCreated.Units {
				roundCreated.Units = append(roundCreated.Units, BattleUnit{
					Team:          unit.Team,
					UnitID:        unit.UnitID,
					UnitType:      unit.UnitType,
					Rank:          unit.Rank,
					MaxHP:         unit.MaxHP,
					HP:            unit.HP,
					MaxMana:       unit.MaxMana,
					Mana:          unit.Mana,
					AttackRate:    unit.AttackRate,
					AttackRange:   unit.AttackRange,
					AttackDamage:  unit.AttackDamage,
					InstantAttack: unit.InstantAttack,
					MoveSpeed:     unit.MoveSpeed,
					Dead:          unit.Dead,
					Placement: Point{
						X: unit.X,
						Y: unit.Y,
					},
				})
			}

			err = ws.WriteJSON(MessageToClient{
				RoundCreated: &roundCreated,
			})

		}

		if event.RoundFinished != nil && event.RoundFinished.PlayerID == playerID {

			roundFinished := RoundFinished{
				PlayerID:         playerID,
				NewCreditsAmount: event.RoundFinished.NewCreditsAmount,
				NewPlayerHealth:  event.RoundFinished.NewPlayerHealth,
				Units:            []Unit{},
			}

			for _, unit := range event.RoundFinished.Units {
				eventUnit := Unit{
					Team:       1,
					UnitID:     unit.UnitID,
					UnitType:   unit.UnitType,
					Tier:       unit.Tier,
					Rank:       unit.Rank,
					HP:         unit.HP,
					Mana:       unit.Mana,
					AttackRate: unit.AttackRate,
				}

				if unit.Placement != nil {
					eventUnit.Placement = &Point{
						X: unit.Placement.X,
						Y: unit.Placement.Y,
					}
				}

				roundFinished.Units = append(roundFinished.Units, eventUnit)
			}

			err = ws.WriteJSON(MessageToClient{
				RoundFinished: &roundFinished,
			})

		}

		if event.PlayerJoined != nil {
			err = ws.WriteJSON(MessageToClient{
				PlayerJoined: &PlayerJoined{
					Player: Player{
						PlayerID: event.PlayerJoined.Player.GetID(),
						Name:     event.PlayerJoined.Player.GetName(),
					},
				},
			})
		}

		if event.PlayerLeft != nil {
			err = ws.WriteJSON(MessageToClient{
				PlayerLeft: &PlayerLeft{
					Player: Player{
						PlayerID: event.PlayerLeft.Player.GetID(),
						Name:     event.PlayerLeft.Player.GetName(),
					},
				},
			})
		}

		if event.BattleUnitHealthChanged != nil && event.BattleUnitHealthChanged.PlayerID == playerID {
			err = ws.WriteJSON(MessageToClient{
				BattleUnitHealthChanged: &BattleUnitHealthChanged{
					PlayerID: playerID,
					UnitID:   event.BattleUnitHealthChanged.UnitID,
					NewHP:    event.BattleUnitHealthChanged.NewHealth,
				},
			})
		}

		if event.BattleUnitManaChanged != nil && event.BattleUnitManaChanged.PlayerID == playerID {
			err = ws.WriteJSON(MessageToClient{
				BattleUnitManaChanged: &BattleUnitManaChanged{
					PlayerID: playerID,
					UnitID:   event.BattleUnitManaChanged.UnitID,
					NewMana:  event.BattleUnitManaChanged.NewMana,
				},
			})
		}

		if err != nil {
			log.Printf("Sending ws message error %v", err)

			return
		}
	}
}
