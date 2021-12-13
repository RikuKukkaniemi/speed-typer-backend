package routes

import (
	"context"
	"log"
	"net/http"
	"rikukukkaniemi/speed-typer-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var wordListCollection *mongo.Collection = OpenCollection(Client, "wordLists")

func AddWordList(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var wordList models.WordList

	if err := c.BindJSON(&wordList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if !IsValid(wordList, c) {
		return
	}

	result, insertErr := wordListCollection.InsertOne(ctx, wordList)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "wordList item was not created"})
		log.Println(insertErr)
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)
}

//get word list by the language
func GetWordListByLanguage(c *gin.Context) {
	language := c.Params.ByName("language")

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var wordList []bson.M

	cursor, err := wordListCollection.Find(ctx, bson.M{"language": language})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if err = cursor.All(ctx, &wordList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, wordList)
}
