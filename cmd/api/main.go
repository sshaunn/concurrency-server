package main

import (
	"os"

	"concurrency-server/config"
	"concurrency-server/internal/common"
	"concurrency-server/internal/middleware"
	routes "concurrency-server/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	route := gin.New()
	route.Use(middleware.Logger())
	routes.Routes(route)
	route.Run(os.Getenv(common.SERVER_PORT))
}
