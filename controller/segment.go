package controller

import (
	"net/http"
	"user-segments/database"
	"user-segments/model"

	"github.com/gin-gonic/gin"
)

func AddSegment(context *gin.Context) {
	var newSegment model.Segment

	if err := context.ShouldBindJSON(&newSegment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedSegment, err := newSegment.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedSegment})
}

func GetAllSegments(context *gin.Context) {
	var segments []model.Segment
	database.Database.Find(&segments)
	for i := range segments {
		var user []*model.User
		database.Database.Model(&segments[i]).Association("Users").Find(&user)
		segments[i].Users = user
	}
	context.JSON(http.StatusOK, gin.H{"data": segments})
}

func GetSegmentByTitle(context *gin.Context) {
	var segment model.Segment
	var user []*model.User
	if err := database.Database.Where("Title = ?", context.Param("title")).First(&segment).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Segment not found!"})
		return
	}
	database.Database.Model(&segment).Association("Users").Find(&user)
	segment.Users = user
	context.JSON(http.StatusOK, gin.H{"data": segment})
}

func DeleteSegmentByTitle(context *gin.Context) {
	var segment model.Segment

	if err := database.Database.Where("Title = ?", context.Param("title")).First(&segment).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Segment not found!"})
		return
	}
	database.Database.Model(&segment).Association("Segments").Clear()
	database.Database.Delete(&segment)

	context.JSON(http.StatusNoContent, gin.H{"data": true})
}
