package routes

import (
	"context"
	"fmt"
	"net/http"
	"rikukukkaniemi/speed-typer-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var highscoreCollection *mongo.Collection = OpenCollection(Client, "highscores")

//add an highscore
func AddHighscore(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var highscore models.Highscore

	if err := c.BindJSON(&highscore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if !IsValid(highscore, c) {
		return
	}
	highscore.ID = primitive.NewObjectID()

	result, insertErr := highscoreCollection.InsertOne(ctx, highscore)
	if insertErr != nil {
		msg := fmt.Sprintf("highscore item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

//get all highscores
func GetHighscores(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var highscores []bson.M

	cursor, err := highscoreCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &highscores); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(highscores)

	c.JSON(http.StatusOK, highscores)
}
