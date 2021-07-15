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
		UnitType:     "unit_droid",
		Rank:         1,
		Tier:         1,
		Cost:         20,
		HP:           100,
		AttackRange:  10,
		AttackRate:   1000,
		AttackDamage: 10,
		MoveSpeed:    20,
	})

	unitPropertyStore.SaveUnit(match.UnitProperties{
		UnitType:     "unit_clone",
		Rank:         1,
		Tier:         1,
		Cost:         65,
		HP:           100,
		AttackRange:  10,
		AttackRate:   100,
		AttackDamage: 10,
		MoveSpeed:    20,
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
