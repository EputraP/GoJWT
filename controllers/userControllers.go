package controllers

import (
	"net/http"

	"github.com/EputraP/GoJWT/initializers"
	"github.com/EputraP/GoJWT/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Username string `gorm:"type:varchar(50);not null"`
		Password string `gorm:"type:varchar(50);not null"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	// create the user
	user := models.UserList{Username: body.Username, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes adding new user",
	})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body

	// look up requested user

	// compare sent in pass with user pass hash

	// generate a jwt token

	// send it back

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
