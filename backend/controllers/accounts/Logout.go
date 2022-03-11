package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
)

/*This handle calls the DeleteSession key function, and if there is no error, it returns a succes message.*/
func Logout(c *gin.Context) {
	if !sessionmanager.DeleteSessionKey(c.Query("token")) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was a problem with logout, please try again and sry for your incovinience.",
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}
