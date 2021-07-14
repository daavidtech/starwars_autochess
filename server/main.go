package main

import (
	"context"

	"github.com/daavidtech/starwars_autochess/game"
	"github.com/daavidtech/starwars_autochess/match"
	"github.com/daavidtech/starwars_autochess/networking"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ctx := context.Background()

	unitPropertyStore := match.NewUnitPropertyStore()

	unitPropertyStore.SaveUnit(match.UnitProperties{
		UnitType: "unit_droid",
		Tier:     1,
		Cost:     20,
	})

	unitPropertyStore.SaveUnit(match.UnitProperties{
		UnitType: "unit_clone",
		Tier:     1,
		Cost:     65,
	})

	tierProbabilities := match.NewTierProbabilities([][]int{
		[]int{
			100,
		},
	})

	gameCoordinator := game.NewGameCoordinator(ctx)

	gameCoordinator.UnitPropertyStore = &unitPropertyStore
	gameCoordinator.TierProbabilities = &tierProbabilities

	userRepo := game.NewUserRepository()

	newUser := game.NewUser()

	newUser.SetUsername("dragonslayer69")
	newUser.SetPassword("asdf")

	userRepo.Save(newUser)

	newUser = game.NewUser()

	newUser.SetUsername("hotmom")
	newUser.SetPassword("asdf")

	userRepo.Save(newUser)

	wsServer := networking.WsServer{
		UserRepository:  userRepo,
		GameCoordinator: gameCoordinator,
	}

	r.GET("/api/socket", wsServer.HandleSocket)

	r.Run(":4100")
}
