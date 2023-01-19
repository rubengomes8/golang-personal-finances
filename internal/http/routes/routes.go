package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/controllers"
)

func SetupRouter(expensesController controllers.ExpensesController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("expense/:id", expensesController.GetExpenseById) // TODO
		v1.POST("expense", expensesController.CreateExpense)
		v1.PUT("expense/:id", expensesController.UpdateExpense)    // TODO
		v1.DELETE("expense/:id", expensesController.DeleteExpense) // TODO
	}
	return r
}

/* CHECKING HERE -> https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8*/
