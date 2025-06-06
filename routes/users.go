package routes

import (
	"events/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user model.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request body"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
