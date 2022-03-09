package accounts

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/controllers"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func Login(c *gin.Context) {
	var userinfo controllers.User
	c.BindJSON(&userinfo)
	//log.Println()
	userinfo.Email = template.HTMLEscapeString(userinfo.Email)
	userinfo.Password = template.HTMLEscapeString(userinfo.Password)
	var user controllers.User
	iter := controllers.Users.Find(bson.M{"email": userinfo.Email}).Iter()
	if !iter.Next(&user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user with this email.",
		})
		c.Abort()
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userinfo.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Password is wrong.",
		})
		c.Abort()
		return
	}
	if user.DeletedAccount != 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account was deleted at: " + fmt.Sprint(user.DeletedAccount),
		})
		c.Abort()
		return
	}
	token := sessionmanager.SetSessionKeys(c.ClientIP(), user.Username)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully", "token": token,
	})
}
