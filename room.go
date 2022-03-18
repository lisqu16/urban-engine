package main

import (
	"encoding/json"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Room struct {
	Subs []*net.Conn
	Msg  chan string
}

var mainRoom = Room{
	make([]*net.Conn, 0),
	make(chan string),
}

func (r *Room) Sub(c *net.Conn) {
	r.Subs = append(r.Subs, c)
}

func init() {
	// room msg handling
	go func() {
		for v := range mainRoom.Msg {
			// create payload
			var rawpld = map[string]interface{}{
				"op":  1,
				"msg": v,
			}
			pld, _ := json.Marshal(rawpld)

			// write payload
			for _, c := range mainRoom.Subs {
				wsutil.WriteServerMessage(*c, ws.OpText, pld)
			}
		}
	}()
}
