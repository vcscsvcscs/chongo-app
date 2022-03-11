package accounts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/controllers"
	"github.com/vcscsvcscs/chongo-app/backend/sessionmanager"
	"gopkg.in/mgo.v2/bson"
)

/*
This handler was made so users can request to delete their data from our database.
This data will be permanently deleted, 15 days after request.
*/
func DeleteAcc(c *gin.Context) {
	token := c.Query("token")
	username := sessionmanager.Users[token]

	if !sessionmanager.DeleteSessionKey(token) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was a problem with deleting your account, please try again and sry for your incovinience.",
		})
		return
	}

	err := controllers.Users.Update(bson.M{"username": username}, bson.M{"$set": bson.M{"deleted": time.Now().Unix()}})
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
