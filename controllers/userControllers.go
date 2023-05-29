package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/EputraP/GoJWT/initializers"
	"github.com/EputraP/GoJWT/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func init() {

}

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
	// Get the username and pass off req body
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
	// look up requested user
	var user models.UserList
	// result := initializers.DB.Where("username = ?", body.Username).First(&body)
	result := initializers.DB.First(&user, "username = ?", body.Username)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get data from DB",
		})
		return
	}
	if user.UserId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username",
		})
		return
	}
	// compare sent in pass with user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})
		return
	}
	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secrete
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
