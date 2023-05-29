package main

import (
	"github.com/EputraP/GoJWT/controllers"
	"github.com/EputraP/GoJWT/initializers"
	"github.com/EputraP/GoJWT/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
