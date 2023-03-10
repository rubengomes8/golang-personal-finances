// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/log_template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package user

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// UserRepoWithLogs implements repository.UserRepo that is instrumented with zerolog logger
type UserRepoWithLogs struct {
	base repository.UserRepo
}

// GetUserByUsername implements repository.UserRepo
func (d UserRepoWithLogs) GetUserByUsername(ctx context.Context, s1 string) (u1 models.UserTable, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"s1":  s1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"u1":  u1,
				"err": err}).Err(err).Str("decorator", "UserRepoWithLogs").Str("method", "GetUserByUsername").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"u1":  u1,
				"err": err}).Str("decorator", "UserRepoWithLogs").Str("method", "GetUserByUsername").Msg("Finish")
		}
	}()
	return d.base.GetUserByUsername(ctx, s1)
}

// InsertUser implements repository.UserRepo
func (d UserRepoWithLogs) InsertUser(ctx context.Context, u1 models.UserTable) (i1 int64, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"u1":  u1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Err(err).Str("decorator", "UserRepoWithLogs").Str("method", "InsertUser").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Str("decorator", "UserRepoWithLogs").Str("method", "InsertUser").Msg("Finish")
		}
	}()
	return d.base.InsertUser(ctx, u1)
}

// NewUserRepoWithLogs instruments an implementation of the repository.UserRepo with simple logging
func NewUserRepoWithLogs(base repository.UserRepo) repository.UserRepo {
	decorate := os.Getenv("DECORATE")
	if decorate == "true" || decorate == "1" {
		return UserRepoWithLogs{
			base: base,
		}
	}

	return base
}
