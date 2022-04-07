package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*This handle calls the DeleteSession key function, and if there is no error, it returns a succes message.*/
func (a *Accounts) Logout(c *gin.Context) {
	if !a.sessionManager.DeleteSessionKey(c.Query("token")) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was a problem with logout, please try again and sry for your incovinience.",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}
