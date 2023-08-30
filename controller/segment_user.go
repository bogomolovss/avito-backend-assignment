package controller

import (
	"fmt"
	"net/http"
	"user-segments/database"
	"user-segments/model"

	"github.com/gin-gonic/gin"
)

func GetActiveUserSegmentsByUsername(context *gin.Context) {
	var user model.User
	var segment []model.Segment
	if err := database.Database.Where("Username= ?", context.Param("username")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	fmt.Println(database.Database.Model(&user).Association("Segments").Find(&segment))
	context.JSON(http.StatusOK, gin.H{"data": segment})
}
