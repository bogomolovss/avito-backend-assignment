package controller

import (
	"net/http"
	"user-segments/database"
	"user-segments/model"

	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {
	var newUser model.User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedSegment, err := newUser.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedSegment})
}

func GetAllUsers(context *gin.Context) {
	var users []model.User
	database.Database.Find(&users)
	for i := range users {
		var segment []model.Segment
		database.Database.Model(&users[i]).Association("Segments").Find(&segment)
		users[i].Segments = segment
	}
	context.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByUsername(context *gin.Context) {
	var user model.User
	var segment []model.Segment
	if err := database.Database.Where("Username= ?", context.Param("username")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	database.Database.Model(&user).Association("Segments").Find(&segment)
	user.Segments = segment
	context.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserByUsername(context *gin.Context) {
	var user model.User
	if err := database.Database.Where("Username= ?", context.Param("username")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	database.Database.Model(&user).Association("Segments").Clear()
	database.Database.Delete(&user)

	context.JSON(http.StatusNoContent, gin.H{"data": true})
}
