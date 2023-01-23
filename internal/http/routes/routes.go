package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/service"
)

// SetupRouter sets up the http routes
func SetupRouter(expensesService service.ExpensesService, authService service.AuthService) *gin.Engine {
	r := gin.Default()

	authentication := r.Group("/auth")
	{
		authentication.POST("register/", authService.Register)
		authentication.POST("login/", authService.Login)
	}

	v1 := r.Group("/v1")
	{
		v1.Use(auth.JwtAuthMiddleware())
		v1.GET("expense/:id", expensesService.GetExpenseByID)
		v1.GET("expenses/dates/:min_date/:max_date", expensesService.GetExpensesByDates)
		v1.GET("expenses/category/:category", expensesService.GetExpensesByCategory)
		v1.GET("expenses/subcategory/:sub_category", expensesService.GetExpensesBySubCategory)
		v1.GET("expenses/card/:card", expensesService.GetExpensesByCard)
		v1.POST("expense", expensesService.CreateExpense)
		v1.PUT("expense/:id", expensesService.UpdateExpense)
		v1.DELETE("expense/:id", expensesService.DeleteExpense)
	}

	return r
}
