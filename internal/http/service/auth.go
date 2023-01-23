package service

import (
	"fmt"
	"net/http"

	"github.com/rubengomes8/golang-personal-finances/internal/http/auth"
	httpModels "github.com/rubengomes8/golang-personal-finances/internal/models/http"
	rdsModels "github.com/rubengomes8/golang-personal-finances/internal/models/rds"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"

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

	var input httpModels.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("wrong body format or incomplete data: %v", err),
		})
		return
	}

	hashedPwd, err := auth.EncryptPassword(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not encrypt user password: %v", err),
		})
		return
	}

	user := rdsModels.UserTable{
		Username: input.Username,
		Passhash: hashedPwd,
		Salt:     "",
	}

	_, err = a.UserRepo.InsertUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("could not insert user: %v", err),
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
	ctx.Writer.Flush()
}

// Login logs in a user
func (a *AuthService) Login(ctx *gin.Context) {

	var input httpModels.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("wrong body format or incomplete data: %v", err),
		})
		return
	}

	userTable, err := a.UserRepo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("error getting user by username: %v", err),
		})
		return
	}

	token, err := auth.LoginCheck(ctx, input.Username, input.Password, userTable)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
			ErrorMsg: fmt.Sprintf("error checking user login: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, httpModels.TokenResponse{
		Token: token,
	})
	ctx.Writer.Flush()
}
