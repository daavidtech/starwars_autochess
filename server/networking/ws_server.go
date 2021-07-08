package networking

import (
	"log"
	"net/http"

	"github.com/daavidtech/starwars_autochess/game"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsServer struct {
	UserRepository  game.UserRepository
	GameCoordinator *game.GameCoordinator
}

func (wsServer *WsServer) HandleSocket(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not websocket handshake", 400)
	} else if err != nil {
		return
	}

	log.Println("New ws connection")

	// ws.WriteJSON(MessageToClient{
	// 	ShopRefilled: &ShopRefilled{
	// 		ShopUnits: []ShopUnit{
	// 			ShopUnit{
	// 				UnitType: "unit_droid",
	// 				Level:    3,
	// 				HP:       200,
	// 				Mana:     300,
	// 				Rank:     3,
	// 				Cost:     100,
	// 			},
	// 		},
	// 	},
	// })

	user := wsServer.UserRepository.Fetch("1")

	currentMatch := user.GetCurrentMatch()

	if currentMatch == nil {
		currentMatch = wsServer.GameCoordinator.FindNewMatch()
		newPlayer := currentMatch.CreatePlayer()

		user.SetCurrentPlayerID(newPlayer.GetID())
		user.SetCurrentMatch(currentMatch)
	}

	matchID := currentMatch.GetID()
	playerID := user.GetCurrentPlayerID()

	// err = ws.WriteJSON(MessageToClient{
	// 	CreateUnit: &CreateUnit{
	// 		ID:       "1",
	// 		UnitType: "UNIT_DROID",
	// 		X:        2,
	// 		Y:        2,
	// 	},
	// })

	// if err != nil {
	// 	log.Println("Failed to send json message", err)
	// }

	// reqCtx := getReqCtx(ctx)

	// playerControls := game.PlayerControls{}

	go func() {
		for {
			var msg MessageFromClient

			err := ws.ReadJSON(&msg)

			if err != nil {
				log.Printf("error while receiving %v", err)

				break
			}

			log.Printf("received msg %v", msg)

			if msg.BuyUnit != nil {
				log.Printf("BuyUnit %v", msg.BuyUnit)

				currentMatch.BuyUnit(playerID, msg.BuyUnit.ShopUnitIndex)
			}

			if msg.PlaceUnit != nil {
				log.Printf("PlaceUnit %v to %v:%v", msg.PlaceUnit.UnitID, msg.PlaceUnit.X, msg.PlaceUnit.Y)

				currentMatch.PlaceUnit(playerID, msg.PlaceUnit.UnitID, msg.PlaceUnit.X, msg.PlaceUnit.Y)
			}

			if msg.SellUnit != nil {
				log.Println("SellUnit")

				currentMatch.SellUnit(playerID, msg.SellUnit.UnitID)
			}

			if msg.BuyLevelUp != nil {
				log.Println("BuyLevelUp")
				currentMatch.BuyLevelUp(playerID)
			}

			if msg.RecycleShopUnits != nil {
				log.Println("RecycleShopUnit")
				currentMatch.RecycleShopUnits(playerID)
			}
		}
	}()

	eventBroker := currentMatch.GetEventBroker()

	ch := eventBroker.Subscribe(matchID)

	for event := range ch {
		// if event.NewBarrackUnit != nil {

		// }

		var err error

		if event.ShopRefilled != nil {
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

		if event.ShopUnitRemoved != nil {
			log.Println("Sending shop unit removed")

			ws.WriteJSON(MessageToClient{
				ShopUnitRemoved: &ShopUnitRemoved{
					ShopUnitID: event.ShopUnitRemoved.ShopUnitID,
				},
			})
		}

		if event.BarrackUnitAdded != nil {
			unitAdded := UnitAdded{
				UnitID:     event.BarrackUnitAdded.UnitID,
				UnitType:   event.BarrackUnitAdded.UnitType,
				Rank:       event.BarrackUnitAdded.Rank,
				HP:         event.BarrackUnitAdded.HP,
				Mana:       event.BarrackUnitAdded.Mana,
				AttackRate: event.BarrackUnitAdded.AttackRate,
			}

			msg := MessageToClient{
				UnitAdded: &unitAdded,
			}

			log.Println("Seding unitAdded to client")

			ws.WriteJSON(msg)
		}

		if event.BarrackUnitRemoved != nil {
			unitRemoved := UnitRemoved{
				UnitID: event.BarrackUnitRemoved.UnitID,
			}

			msg := MessageToClient{
				UnitRemoved: &unitRemoved,
			}

			log.Println("Sending unitRemoved to client")

			ws.WriteJSON(msg)
		}

		if event.BarrackUnitUpgraded != nil {
			unitUpgraded := UnitUpgraded{
				UnitID:     event.BarrackUnitUpgraded.UnitID,
				Rank:       event.BarrackUnitUpgraded.Rank,
				HP:         event.BarrackUnitUpgraded.HP,
				Mana:       event.BarrackUnitUpgraded.Mana,
				AttackRate: event.BarrackUnitUpgraded.AttackRate,
			}

			log.Println("Sending unitUpgraded to client")

			ws.WriteJSON(MessageToClient{
				UnitUpgraded: &unitUpgraded,
			})
		}

		if event.UnitPlaced != nil {
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

			err = ws.WriteJSON(MessageToClient{
				MatchPhaseChanged: &MatchPhaseChanged{
					MatchPhase: event.PhaseChanged.MatchPhase,
				},
			})
		}

		if event.UnitStartedMovingTo != nil {
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

		if event.UnitArrivedTo != nil {
			log.Printf("Sending unit %v arrived to", event.UnitArrivedTo.UnitID)

			err = ws.WriteJSON(MessageToClient{
				UnitArrivedTo: &UnitArrivedTo{
					UnitID: event.UnitArrivedTo.UnitID,
					X:      event.UnitArrivedTo.X,
					Y:      event.UnitArrivedTo.Y,
				},
			})
		}

		if event.UnitDied != nil {
			log.Printf("Sending unit %v died to client", event.UnitDied.UnitID)

			err = ws.WriteJSON(MessageToClient{
				UnitDied: &UnitDied{
					UnitID: event.UnitDied.UnitID,
				},
			})
		}

		if event.RoundCreated != nil {
			units := []BattleUnit{}

			for _, unit := range event.RoundCreated.Units {
				units = append(units, BattleUnit{
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
					X:             int(unit.X),
					Y:             int(unit.Y),
				})
			}

			err = ws.WriteJSON(MessageToClient{
				RoundCreated: &RoundCreated{
					Units: units,
				},
			})
		}

		if err != nil {
			log.Printf("Sending ws message error %v", err)

			return
		}
	}

	log.Println("Event broker subscription stopped")
}
