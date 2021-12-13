package routes

import (
	"log"
	"net/http"
	"rikukukkaniemi/speed-typer-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func IsValid(toBeValidatable models.Validatable, c *gin.Context) bool {
	validationErr := validate.Struct(toBeValidatable)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		log.Println(validationErr)
		return false
	}

	return true
}
