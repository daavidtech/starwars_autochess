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

	ws.WriteJSON(MessageToClient{
		GamePhaseChanged: &GamePhaseChanged{
			LoginPhase,
		},
	})

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
				GamePhaseChanged: &GamePhaseChanged{
					GamePhase: MainMenuPhase,
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

	if currentMatch == nil || currentMatch.GetMatchPhase() == match.EndPhase {
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

	go handle_match_events(ws, currentMatch, user.GetCurrentPlayerID())

	currentMatch.StartLobby()

	for {
		var msg MessageFromClient

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error while receiving %v", err)

			break
		}

		if msg.StartMatch != nil {
			currentMatch.Start(user.GetCurrentPlayerID())

			continue
		}

		if msg.BuyUnit != nil {
			log.Printf("BuyUnit %v", msg.BuyUnit)

			currentMatch.BuyUnit(user.GetCurrentPlayerID(), msg.BuyUnit.ShopUnitIndex)
		}

		if msg.PlaceUnit != nil {
			log.Printf("PlaceUnit %v to %v:%v", msg.PlaceUnit.UnitID, msg.PlaceUnit.X, msg.PlaceUnit.Y)

			currentMatch.PlaceUnit(user.GetCurrentPlayerID(), msg.PlaceUnit.UnitID, msg.PlaceUnit.X, msg.PlaceUnit.Y)
		}

		if msg.SellUnit != nil {
			log.Println("SellUnit")

			currentMatch.SellUnit(user.GetCurrentPlayerID(), msg.SellUnit.UnitID)
		}

		if msg.BuyLevelUp != nil {
			log.Println("BuyLevelUp")
			currentMatch.BuyLevelUp(user.GetCurrentPlayerID())
		}

		if msg.RecycleShopUnits != nil {
			log.Println("RecycleShopUnit")
			currentMatch.RecycleShopUnits(user.GetCurrentPlayerID())
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

			go handle_match_events(ws, currentMatch, user.GetCurrentPlayerID())
		}
	}
}
