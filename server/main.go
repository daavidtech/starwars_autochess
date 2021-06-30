package main

import (
	"context"

	"github.com/daavidtech/starwars_autochess/game"
	"github.com/daavidtech/starwars_autochess/networking"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ctx := context.Background()

	gameCoordinator := game.NewGameCoordinator(ctx)

	userRepo := game.NewUserRepository()

	wsServer := networking.WsServer{
		UserRepository:  userRepo,
		GameCoordinator: gameCoordinator,
	}

	r.GET("/api/socket", wsServer.HandleSocket)

	r.Run(":4100")
}
