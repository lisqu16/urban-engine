package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Static("/static", "./static/")

	r.GET("/gateway", gin.WrapF(wse))

	log.Fatalln(r.Run(":8080"))
}
