package main

import (
	"log"

	"github.com/bit365/iotcloud/services/webapi/config"
	"github.com/bit365/iotcloud/services/webapi/internal/database"
	"github.com/bit365/iotcloud/services/webapi/internal/middleware"
	"github.com/bit365/iotcloud/services/webapi/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	appConfig := config.LoadConfig()

	router := gin.Default()

	database_connection_string := appConfig.GetString("database.iotcloud")
	database.RunMigrations(database_connection_string)
	db, err := database.InitDB(database_connection_string)

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Use(middleware.Logger())
	routers.SetupRouter(router, db)

	addr := appConfig.GetString("server.address")
	router.Run(addr)
}
