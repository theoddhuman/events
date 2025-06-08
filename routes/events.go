package routes

import (
	"events/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func getEvents(context *gin.Context) {
	events, err := model.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request body"})
		return
	}
	event.UserId = context.GetInt64("userId")
	event.DateTime = time.Now()
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse updatedEven id"})
		return
	}

	event, err := model.GetEventById(eventId)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}

	var updatedEvent model.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request body"})
		return
	}
	updatedEvent.Id = eventId
	updatedEvent.UserId = userId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse updatedEven id"})
		return
	}

	event, err := model.GetEventById(eventId)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User does not have permission"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse event id"})
		return
	}
	event, err := model.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	}
	context.JSON(http.StatusOK, event)
}
