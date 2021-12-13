package routes

import (
	"context"
	"log"
	"net/http"
	"rikukukkaniemi/speed-typer-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var highscoreCollection *mongo.Collection = OpenCollection(Client, "highscores")

//add an highscore
func AddHighscore(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var highscore models.Highscore

	if err := c.BindJSON(&highscore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if !IsValid(highscore, c) {
		return
	}

	highscore.ID = primitive.NewObjectID()
	highscore.Date = primitive.NewDateTimeFromTime(time.Now())

	_, insertErr := highscoreCollection.InsertOne(ctx, highscore)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "highscore item was not created"})
		log.Println(insertErr)
		return
	}
	defer cancel()

	GetHighscores(c)
}

//get top 10 highscores
func GetHighscores(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var highscores []bson.M

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"score", -1}})
	findOptions.SetLimit(10)

	cursor, err := highscoreCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if err = cursor.All(ctx, &highscores); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, highscores)
}
