package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/instrumentation"
)

// SetupRouter sets up the http routes
func SetupRouter(
	expensesHandlers handlers.Expenses,
	incomesHandlers handlers.Incomes,
	authHandlers handlers.Auth,
) *gin.Engine {

	r := gin.Default()
	r.Handle(http.MethodGet, "/metrics", gin.WrapH(instrumentation.RegistryHandler()))

	authentication := r.Group("/auth")
	{
		authentication.POST("register/", authHandlers.Register)
		authentication.POST("login/", authHandlers.Login)
	}

	v1 := r.Group("/v1")
	{
		// Use authentication
		v1.Use(auth.JwtAuthMiddleware())

		// Expenses
		v1.GET("expense/:id", expensesHandlers.GetExpenseByID)
		v1.POST("expense", expensesHandlers.CreateExpense)
		v1.PUT("expense/:id", expensesHandlers.UpdateExpense)
		v1.DELETE("expense/:id", expensesHandlers.DeleteExpense)
		v1.GET("expenses/dates/:min_date/:max_date", expensesHandlers.GetExpensesByDates)
		v1.GET("expenses/category/:category", expensesHandlers.GetExpensesByCategory)
		v1.GET("expenses/subcategory/:sub_category", expensesHandlers.GetExpensesBySubCategory)
		v1.GET("expenses/card/:card", expensesHandlers.GetExpensesByCard)

		// Incomes
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
