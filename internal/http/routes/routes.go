package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
)

// SetupRouter sets up the http routes
func SetupRouter(
	expensesHandler handlers.Expenses,
	incomesHandler handlers.Incomes,
	authHandler handlers.AuthService,
) *gin.Engine {
	r := gin.Default()

	authentication := r.Group("/auth")
	{
		authentication.POST("register/", authHandler.Register)
		authentication.POST("login/", authHandler.Login)
	}

	v1 := r.Group("/v1")
	{
		v1.Use(auth.JwtAuthMiddleware())

		v1.GET("expense/:id", expensesHandler.GetExpenseByID)
		v1.POST("expense", expensesHandler.CreateExpense)
		v1.PUT("expense/:id", expensesHandler.UpdateExpense)
		v1.DELETE("expense/:id", expensesHandler.DeleteExpense)
		v1.GET("expenses/dates/:min_date/:max_date", expensesHandler.GetExpensesByDates)
		v1.GET("expenses/category/:category", expensesHandler.GetExpensesByCategory)
		v1.GET("expenses/subcategory/:sub_category", expensesHandler.GetExpensesBySubCategory)
		v1.GET("expenses/card/:card", expensesHandler.GetExpensesByCard)

		v1.GET("income/:id", incomesHandler.HandleGetIncomeByID)
		v1.POST("income", incomesHandler.HandleCreateIncome)
		v1.PUT("income/:id", incomesHandler.HandleUpdateIncome)
		v1.DELETE("income/:id", incomesHandler.HandleDeleteIncome)
		v1.GET("incomes/category/:category", incomesHandler.HandleGetIncomesByCategory)
		v1.GET("incomes/card/:card", incomesHandler.HandleGetIncomesByCard)
		v1.GET("incomes/dates/:min_date/:max_date", incomesHandler.HandleGetIncomesByDates)
	}

	return r
}
