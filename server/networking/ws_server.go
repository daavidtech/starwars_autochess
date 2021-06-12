package networking

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsServer struct {
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

	x := 0

	ws.WriteJSON(MessageToClient{
		CreateUnit: &CreateUnit{
			ID:       "1",
			UnitType: "UNIT_DROID",
			X:        x,
			Y:        2,
		},
	})

	for {
		ws.WriteJSON(MessageToClient{
			ChangeUnitPosition: &ChangeUnitPosition{
				ID: "1",
				X:  x,
			},
		})

		timer := time.NewTimer(20 * time.Millisecond)
		<-timer.C

		x += 2
	}

	// for i := 0; i < 10; i++ {
	// 	ws.WriteJSON(MessageToClient{
	// 		CreateUnit: &CreateUnit{
	// 			ID:       "1",
	// 			UnitType: "UNIT_DROID",
	// 			X:        20 * i,
	// 			Y:        2,
	// 		},
	// 	})
	// }
}
