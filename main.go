package main

import (
	"dream_11/database"
	"dream_11/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	router.RegisterRouter(r)
	
	fmt.Println("Server starts at 8080.........")
	r.Run(":8080")
}
