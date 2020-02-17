package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "application/json", []byte(`{"serverTime": `+time.Now().Format(time.RFC3339)+`}`))
	})

	log.Fatal(router.Run(":8080"))
}
