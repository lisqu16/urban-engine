package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func wse(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Fatalln(err)
	}

	// subscribe main room
	mainRoom.Sub(&conn)

	go func() {
		defer conn.Close()

		// reading data
		for {
			rawmsg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Fatalln(err)
			}

			// json
			var msg map[string]interface{}
			if err := json.Unmarshal(rawmsg, &msg); err != nil {
				log.Fatalln(err)
			}

			// handling msg
			switch msg["op"].(float64) {
			// case 0: // todo: authenticating
			case 1: // writing message
				mainRoom.Msg <- msg["msg"].(string)
			}
		}
	}()
}
