package accounts

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/controllers"
	"github.com/vcscsvcscs/chongo-app/backend/utilities"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//This function registers a new document in the database with the new users data. It can return multiple types of error messages or at succesful registration session token.
func (a *Accounts) Register(c *gin.Context) {
	var userinfo controllers.User
	log.Println(c.BindJSON(&userinfo))
	userinfo.Name = template.HTMLEscapeString(userinfo.Name)
	userinfo.Username = template.HTMLEscapeString(userinfo.Username)
	userinfo.Password = template.HTMLEscapeString(userinfo.Password)
	if !utilities.IsEmailValid(userinfo.Email) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email is not valid.",
		})
		c.Abort()
		return
	}
	userinfo.Email = template.HTMLEscapeString(userinfo.Email)
	var user controllers.User
	iter := controllers.Users.Find(bson.M{"email": userinfo.Email}).Iter()
	if iter.Next(&user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There is already a user with this email.",
		})
		c.Abort()
		return
	}
	iter = controllers.Users.Find(bson.M{"username": userinfo.Username}).Iter()
	if iter.Next(&user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There is already a user with this username.",
		})
		c.Abort()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userinfo.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "This password is not supported.",
		})
		log.Println(err)
		c.Abort()
		return
	}
	userinfo.Password = string(hash)
	userinfo.Ip = c.ClientIP()
	userinfo.DeletedAccount = 0
	if err := controllers.Users.Insert(&userinfo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was an error with our database, please try again, sry for the incovinience.",
		})
		log.Println(err)
		c.Abort()
		return
	}
	token := a.sessionManager.SetSessionKeys(c.ClientIP(), userinfo.Username)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registration successfull", "token": token,
	})
}
