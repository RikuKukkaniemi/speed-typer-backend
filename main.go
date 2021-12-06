package main

import (
	"log"
	"os"
	"rikukukkaniemi/speed-typer-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/highscores", routes.GetHighscores)
	router.POST("/highscores", routes.AddHighscore)
	router.Run(":" + os.Getenv("PORT"))
}
