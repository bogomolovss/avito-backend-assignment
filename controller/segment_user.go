package controller

import (
	"net/http"
	"user-segments/database"
	"user-segments/model"

	"github.com/gin-gonic/gin"
)

// func AddUserToSegment(context *gin.Context) {
// 	var user model.User
// 	var segment []model.Segment
// 	if err := database.Database.Where("Username= ?", context.Param("username")).First(&user).Error; err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
// 		return
// 	}
// 	database.Database.Model(&user).Association("Segments").Find(&segment)
// 	context.JSON(http.StatusOK, gin.H{"data": segment})
// }

type SegmentLite struct {
	Title string
}

func GetActiveUserSegmentsByUsername(context *gin.Context) {
	var user model.User
	var segment []model.Segment
	if err := database.Database.Where("id= ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	database.Database.Model(&user).Association("Segments").Find(&segment)
	context.JSON(http.StatusOK, gin.H{"data": segment})
}

func GetActiveUserSegmentsByUsernameLite(context *gin.Context) {
	var user model.User
	var segment []model.Segment
	var segments_lite []SegmentLite
	if err := database.Database.Where("id= ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	database.Database.Model(&user).Association("Segments").Find(&segment)
	for i := range segment {
		s := SegmentLite{
			Title: segment[i].Title,
		}
		segments_lite = append(segments_lite, s)
	}
	context.JSON(http.StatusOK, gin.H{"data": segments_lite})
}
