package main

import (
	"events/db"
	"events/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := model.GetEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request body"})
		return
	}
	event.Id = 1
	event.UserId = 1
	event.DateTime = time.Now()
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
