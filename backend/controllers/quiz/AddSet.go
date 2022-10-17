package quiz

import (
	"crypto/md5"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcscsvcscs/chongo-app/backend/controllers/quiz/model"
)

//This function Creates a set of questions in the database with the user provided data. It can return multiple types of error messages or the Quiz id.
func (a *Quizes) AddSet(c *gin.Context) {
	var quiz model.Quiz
	_ = c.BindJSON(&quiz)
	quiz.Name = template.HTMLEscapeString(quiz.Name)
	quiz.Owner = template.HTMLEscapeString(quiz.Owner)
	quiz.Id = string(md5.New().Sum([]byte(quiz.Name + quiz.Owner + time.Now().String())))
	if err := a.db.Insert(&quiz); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "There was an error with our database, please try again, sry for the incovinience.",
		})
		log.Println(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Quiz succesfully created", "id": quiz.Id,
	})
}
