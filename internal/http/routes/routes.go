package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/rubengomes8/golang-personal-finances/internal/http/service"
)

// SetupRouter sets up the http routes
func SetupRouter(expensesController controllers.ExpensesService) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("expense/:id", expensesController.GetExpenseByID)
		v1.GET("expenses/dates/:min_date/:max_date", expensesController.GetExpensesByDates)
		v1.GET("expenses/category/:category", expensesController.GetExpensesByCategory)
		v1.GET("expenses/subcategory/:sub_category", expensesController.GetExpensesBySubCategory)
		v1.GET("expenses/card/:card", expensesController.GetExpensesByCard)
		v1.POST("expense", expensesController.CreateExpense)
		v1.PUT("expense/:id", expensesController.UpdateExpense)
		v1.DELETE("expense/:id", expensesController.DeleteExpense)
	}

	return r
}
