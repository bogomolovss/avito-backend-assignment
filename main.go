package main

import (
	"log"
	"user-segments/controller"
	"user-segments/database"
	"user-segments/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Segment{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/api")
	publicRoutes.POST("/segments", controller.AddSegment)
	publicRoutes.GET("/segments", controller.GetAllSegments)
	publicRoutes.GET("/segments/:title", controller.GetSegmentByTitle)
	publicRoutes.DELETE("/segments/:title", controller.DeleteSegmentByTitle)
	//user part
	publicRoutes.POST("/users", controller.AddUser)
	publicRoutes.GET("/users", controller.GetAllUsers)
	publicRoutes.GET("/users/:username", controller.GetUserByUsername)
	publicRoutes.DELETE("/users/:username", controller.DeleteUserByUsername)
	//user_segment part
	publicRoutes.GET("/user_segments/:id", controller.GetActiveUserSegmentsByUsername)
	router.Run()
}

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}
