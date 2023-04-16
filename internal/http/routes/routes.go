package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/handlers"
	"github.com/rubengomes8/golang-personal-finances/internal/instrumentation"

	"github.com/rubengomes8/golang-personal-finances/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// SetupRouter sets up the http routes
func SetupRouter(
	expensesHandlers handlers.Expenses,
	incomesHandlers handlers.Incomes,
	authHandlers handlers.Auth,
) *gin.Engine {

	r := gin.Default()
	r.Handle(http.MethodGet, "/metrics", gin.WrapH(instrumentation.RegistryHandler()))

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Finances API"
	docs.SwaggerInfo.Description = "This is a REST API to CRUD expenses and incomes."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
