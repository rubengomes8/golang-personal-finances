package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
)

// SetupRouter sets up the http routes
func SetupRouter(
	expensesHandlers handlers.Expenses,
	incomesHandlers handlers.Incomes,
	authHandlers handlers.Auth,
) *gin.Engine {
	r := gin.Default()

	authentication := r.Group("/auth")
	{
		authentication.POST("register/", authHandlers.Register)
		authentication.POST("login/", authHandlers.Login)
	}

	v1 := r.Group("/v1")
	{
		v1.Use(auth.JwtAuthMiddleware())

		v1.GET("expense/:id", expensesHandlers.GetExpenseByID)
		v1.POST("expense", expensesHandlers.CreateExpense)
		v1.PUT("expense/:id", expensesHandlers.UpdateExpense)
		v1.DELETE("expense/:id", expensesHandlers.DeleteExpense)
		v1.GET("expenses/dates/:min_date/:max_date", expensesHandlers.GetExpensesByDates)
		v1.GET("expenses/category/:category", expensesHandlers.GetExpensesByCategory)
		v1.GET("expenses/subcategory/:sub_category", expensesHandlers.GetExpensesBySubCategory)
		v1.GET("expenses/card/:card", expensesHandlers.GetExpensesByCard)

		v1.GET("income/:id", incomesHandlers.HandleGetByID)
		v1.POST("income", incomesHandlers.HandleCreateIncome)
		v1.PUT("income/:id", incomesHandlers.HandleUpdateIncome)
		v1.DELETE("income/:id", incomesHandlers.HandleDeleteIncome)
		v1.GET("incomes/category/:category", incomesHandlers.HandleGetIncomesByCategory)
		v1.GET("incomes/card/:card", incomesHandlers.HandleGetIncomesByCard)
		v1.GET("incomes/dates/:min_date/:max_date", incomesHandlers.HandleGetIncomesByDates)
	}

	return r
}
