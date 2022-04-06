package accounts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
This handler was made so users can request to delete their data from our database.
This data will be permanently deleted, 15 days after request.
*/
func (a *Accounts) DeleteAcc(c *gin.Context) {
	token := c.Query("token")
	username := a.sessionManager.GetUser(token)

	if !a.sessionManager.DeleteSessionKey(token) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was a problem with deleting your account, please try again and sry for your incovinience.",
		})
		return
	}

	err := a.db.Update(username, time.Now())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was a problem with deleting your account, please try again and sry for your incovinience.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User account deleted succesfully, we delete every data permanently after 2 weeks of calming period",
	})
}
