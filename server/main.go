package main

import (
	"github.com/daavidtech/starwars_autochess/networking"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	wsServer := networking.WsServer{}

	r.GET("/api/socket", wsServer.HandleSocket)

	r.Run(":4100")
}
