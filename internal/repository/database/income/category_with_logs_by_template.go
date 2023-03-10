// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/log_template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package income

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// IncomeCategoryRepoWithLogs implements repository.IncomeCategoryRepo that is instrumented with zerolog logger
type IncomeCategoryRepoWithLogs struct {
	base repository.IncomeCategoryRepo
}

// DeleteIncomeCategory implements repository.IncomeCategoryRepo
func (d IncomeCategoryRepoWithLogs) DeleteIncomeCategory(ctx context.Context, i1 int64) (err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"err": err}).Err(err).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "DeleteIncomeCategory").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"err": err}).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "DeleteIncomeCategory").Msg("Finish")
		}
	}()
	return d.base.DeleteIncomeCategory(ctx, i1)
}

// GetIncomeCategoryByID implements repository.IncomeCategoryRepo
func (d IncomeCategoryRepoWithLogs) GetIncomeCategoryByID(ctx context.Context, i1 int64) (i2 models.IncomeCategoryTable, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Err(err).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "GetIncomeCategoryByID").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "GetIncomeCategoryByID").Msg("Finish")
		}
	}()
	return d.base.GetIncomeCategoryByID(ctx, i1)
}

// GetIncomeCategoryByName implements repository.IncomeCategoryRepo
func (d IncomeCategoryRepoWithLogs) GetIncomeCategoryByName(ctx context.Context, s1 string) (i1 models.IncomeCategoryTable, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"s1":  s1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Err(err).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "GetIncomeCategoryByName").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "GetIncomeCategoryByName").Msg("Finish")
		}
	}()
	return d.base.GetIncomeCategoryByName(ctx, s1)
}

// InsertIncomeCategory implements repository.IncomeCategoryRepo
func (d IncomeCategoryRepoWithLogs) InsertIncomeCategory(ctx context.Context, i1 models.IncomeCategoryTable) (i2 int64, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Err(err).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "InsertIncomeCategory").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "InsertIncomeCategory").Msg("Finish")
		}
	}()
	return d.base.InsertIncomeCategory(ctx, i1)
}

// UpdateIncomeCategory implements repository.IncomeCategoryRepo
func (d IncomeCategoryRepoWithLogs) UpdateIncomeCategory(ctx context.Context, i1 models.IncomeCategoryTable) (i2 int64, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Err(err).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "UpdateIncomeCategory").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i2":  i2,
				"err": err}).Str("decorator", "IncomeCategoryRepoWithLogs").Str("method", "UpdateIncomeCategory").Msg("Finish")
		}
	}()
	return d.base.UpdateIncomeCategory(ctx, i1)
}

// NewIncomeCategoryRepoWithLogs instruments an implementation of the repository.IncomeCategoryRepo with simple logging
func NewIncomeCategoryRepoWithLogs(base repository.IncomeCategoryRepo) repository.IncomeCategoryRepo {
	decorate := os.Getenv("DECORATE")
	if decorate == "true" || decorate == "1" {
		return IncomeCategoryRepoWithLogs{
			base: base,
		}
	}

	return base
}
