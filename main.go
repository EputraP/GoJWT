package main

import (
	"github.com/EputraP/GoJWT/controllers"
	"github.com/EputraP/GoJWT/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.GET("/login", controllers.Login)
	r.Run()
}
