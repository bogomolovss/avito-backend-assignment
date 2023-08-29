package main

import (
	"log"
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

func main() {
	router := gin.Default()

	loadEnv()
	loadDatabase()
}
