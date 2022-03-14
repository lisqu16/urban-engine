package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func main() {
	r := gin.New()

	r.Static("/static", "./static/")

	r.Any("/ws", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Fatalln(err)
		}
		defer c.Close(websocket.StatusInternalError, "dupa")
	}))

	log.Fatalln(r.Run(":8080"))
}
