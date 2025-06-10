package routes

import (
	"log"

	"github.com/RyuichiroYoshida/quest-board-project/di"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, c *di.Container) {
	authH := c.AuthHandler
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/login/discord", authH.LoginDiscord)
			auth.GET("/exchange", authH.ExchangeCode)
			auth.GET("/me", authH.Me)
			auth.GET("/logout", authH.Logout)
		}
	}
	log.Println(api)
}
