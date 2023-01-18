package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/rubengomes8/golang-personal-finances/internal/models/client"
)

func CreateExpense(context *gin.Context) {
	// TODO

	var expense models.Expense
	err := json.NewDecoder(context.Request.Body).Decode(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"expense": expense})

}

func UpdateExpense(context *gin.Context) {
	// TODO
}

func GetExpense(context *gin.Context) {
	// TODO
}

func DeleteExpense(context *gin.Context) {
	// TODO
}
