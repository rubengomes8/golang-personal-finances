package service

import (
	"log"
	"net/http"

	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	"github.com/rubengomes8/golang-personal-finances/internal/http/models"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	dbModels "github.com/rubengomes8/golang-personal-finances/internal/repository/models"

	"github.com/gin-gonic/gin"
)

// AuthService handles the authentication requests
type AuthService struct {
	UserRepo repository.UserRepo
}

// NewAuthService creates a new AuthService
func NewAuthService(userRepo repository.UserRepo) (AuthService, error) {
	return AuthService{
		UserRepo: userRepo,
	}, nil
}

// Register registers a user on the database
func (a *AuthService) Register(ctx *gin.Context) {

	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Printf("could not bind register json: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "wrong body format or incomplete data",
		})
		return
	}

	hashedPwd, err := auth.EncryptPassword(input.Username, input.Password)
	if err != nil {
		log.Printf("could not encrypt user password: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorMsg: "could not encrypt user password",
		})
		return
	}

	user := dbModels.UserTable{
		Username: input.Username,
		Passhash: hashedPwd,
	}

	_, err = a.UserRepo.InsertUser(ctx, user)
	if err != nil {
		log.Printf("could not insert user: %v", user)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not create user",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// Login logs in a user
func (a *AuthService) Login(ctx *gin.Context) {

	var input models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Printf("could not bind login json: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "wrong body format or incomplete data",
		})
		return
	}

	userTable, err := a.UserRepo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		log.Printf("error getting user by username: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not get user. Maybe he does not exist",
		})
		return
	}

	token, err := auth.LoginCheck(ctx, input.Username, input.Password, userTable)
	if err != nil {
		log.Printf("error validating login credentials: %v", err)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorMsg: "could not login user",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.TokenResponse{
		Token: token,
	})
	ctx.Writer.Flush()
}
