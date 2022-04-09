package accounts

import (
	"log"
	"net/http"
	"text/template"

	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/model"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/utilities"
	"golang.org/x/crypto/bcrypt"
)

//This function registers a new document in the database with the new users data. It can return multiple types of error messages or at succesful registration session token.
func (a *Accounts) Register(c *gin.Context) {
	var userinfo model.User
	//log.Println(c.BindJSON(&userinfo))
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
	var user model.User
	if a.db.FindByEmail(userinfo.Email, &user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There is already a user with this email.",
		})
		c.Abort()
		return
	}

	if a.db.FindByUserName(userinfo.Username, &user) {
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
	if err := a.db.Insert(&userinfo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was an error with our database, please try again, sry for the incovinience.",
		})
		log.Println(err)
		c.Abort()
		return
	}
	token, err := a.sessionManager.SetSessionKeys(c.ClientIP(), userinfo.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registration successfull", "token": token,
	})
}
