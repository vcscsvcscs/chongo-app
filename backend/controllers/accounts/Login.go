package accounts

import (
	"fmt"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/accounts/model"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*A dead simple Login api, which returns an error message or a session token, which is set by the session manager. */
func (a *Accounts) Login(c *gin.Context) {
	var userinfo model.User
	_ = c.BindJSON(&userinfo)
	//log.Println()
	userinfo.Email = template.HTMLEscapeString(userinfo.Email)
	userinfo.Password = template.HTMLEscapeString(userinfo.Password)
	var user model.User

	if !a.db.FindByEmail(userinfo.Email, &user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user with this email.",
		})
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
	token, err := a.sessionManager.SetSessionKeys(c.ClientIP(), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully", "token": token,
	})
}
