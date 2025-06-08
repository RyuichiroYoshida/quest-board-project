package main

import (
	"github.com/RyuichiroYoshida/quest-board-project/db"
	"github.com/RyuichiroYoshida/quest-board-project/di"
	"github.com/RyuichiroYoshida/quest-board-project/routes"
	"github.com/RyuichiroYoshida/quest-board-project/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.SetupDb()

	container := di.InitContainer(database)

	router := gin.Default()
	routes.Setup(router, container)

	if err := router.Run(":8080"); err != nil {
		utils.LogError("Failed to start server", err)
	} else {
		utils.LogInfo("Server started successfully on port 8080")
	}
}
