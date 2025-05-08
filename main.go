package main

import (
	"go-dev/module/models"
	"net/http"
	"unicode"
	"time"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	server.GET("/",hello)
	server.GET("/events", getAllEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func getAllEvents(c *gin.Context) {
	events := models.GetEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	c.BindJSON(&event)
	event.ID = len(models.GetEvents())
	event.UserID = unicode.ASCII_Hex_Digit.LatinOffset
	event.DateTime = time.Now()
	event.Save()
	c.JSON(http.StatusOK, event)
}