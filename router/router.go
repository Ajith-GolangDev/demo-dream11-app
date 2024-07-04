package router

import (
	"dream_11/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// API's for admin
	r.POST("/create-contest", controllers.CreateContest)
	r.POST("/create-player", controllers.CreatePlayer)

	// API's for users
	r.POST("/loadmoney/:user_id/:amount", controllers.LoadMoney)
	r.POST("/joincontest/:user_id/:contest_id", controllers.JoinContest)
	r.POST("/createteam", controllers.CreateTeam)
	r.GET("/viewteam/:team_id", controllers.ViewTeam)
}
