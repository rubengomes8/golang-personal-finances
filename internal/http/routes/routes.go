package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("expense/:id", controllers.GetExpense)
		v1.POST("expense", controllers.CreateExpense)
		v1.PUT("expense/:id", controllers.UpdateExpense)
		v1.DELETE("expense/:id", controllers.DeleteExpense)
	}
	return r
}

/* CHECKING HERE -> https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8*/
