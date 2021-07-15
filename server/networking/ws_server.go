package networking

import (
	"log"
	"net/http"

	"github.com/daavidtech/starwars_autochess/game"
	"github.com/daavidtech/starwars_autochess/match"
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

	var user *game.User

	for {
		var msg MessageFromClient

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("Error while receiving message %v", err)

			return
		}

		if msg.Login != nil {
			log.Printf("Login attempt %v", msg.Login.Username)

			user = wsServer.UserRepository.FetchByUsername(msg.Login.Username)

			if user == nil {
				log.Printf("User not found %v", msg.Login.Username)
				continue
			}

			if user.GetPassword() != msg.Login.Password {
				log.Printf("Password is incorrect for user %v", msg.Login.Username)

				continue
			}

			log.Printf("Login successfull %v", msg.Login.Username)

			err = ws.WriteJSON(MessageToClient{
				LoginSuccess: &LoginSuccess{
					Uusername: msg.Login.Username,
				},
			})

			if err != nil {
				log.Println("Failed sending loging success to client")
			}

			break
		}
	}

	matchID := user.GetCurrentMatchID()

	var currentMatch *match.Match

	currentMatch = wsServer.GameCoordinator.FindMatch(matchID)

	if currentMatch == nil {
		for {
			var msg MessageFromClient

			err := ws.ReadJSON(&msg)

			if err != nil {
				log.Printf("Error while receiving message %v", err)

				return
			}

			if msg.FindMatch != nil {
				log.Println("Find new match")

				currentMatch = wsServer.GameCoordinator.FindNewMatch()

				matchID = currentMatch.GetID()

				lobbyAdmin := false

				if currentMatch.CountPlayers() == 0 {
					lobbyAdmin = true
				}

				newPlayer := currentMatch.CreatePlayer(user.GetUsername(), lobbyAdmin)

				user.SetCurrentPlayerID(newPlayer.GetID())
				user.SetCurrentMatchID(currentMatch.GetID())

				break
			}
		}
	}

	log.Printf("Current match %v", matchID)

	players := []Player{}

	for _, p := range currentMatch.GetPlayers() {
		players = append(players, Player{
			PlayerID: p.GetID(),
			Name:     p.GetName(),
		})
	}

	ws.WriteJSON(MessageToClient{
		CurrentMatch: &CurrentMatch{
			MatchID: matchID,
			Phase:   currentMatch.GetMatchPhase(),
			Players: players,
		},
	})

	playerID := user.GetCurrentPlayerID()

	go func() {
		eventBroker := currentMatch.GetEventBroker()

		ch := eventBroker.Subscribe(matchID)

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

				err = ws.WriteJSON(MessageToClient{
					MatchPhaseChanged: &MatchPhaseChanged{
						MatchPhase: event.PhaseChanged.MatchPhase,
					},
				})
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

			if err != nil {
				eventBroker.Unsubscribe(ch)

				log.Printf("Sending ws message error %v", err)

				return
			}
		}
	}()

	currentMatch.StartLobby()

	matchID = currentMatch.GetID()

	for {
		var msg MessageFromClient

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error while receiving %v", err)

			break
		}

		if msg.StartMatch != nil {
			currentMatch.Start(playerID)

			continue
		}

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
}
