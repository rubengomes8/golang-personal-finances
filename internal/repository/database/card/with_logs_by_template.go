// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/log_template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package card

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

// CardRepoWithLogs implements repository.CardRepo that is instrumented with zerolog logger
type CardRepoWithLogs struct {
	base repository.CardRepo
}

// DeleteCard implements repository.CardRepo
func (d CardRepoWithLogs) DeleteCard(ctx context.Context, i1 int64) (err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"err": err}).Err(err).Str("decorator", "CardRepoWithLogs").Str("method", "DeleteCard").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"err": err}).Str("decorator", "CardRepoWithLogs").Str("method", "DeleteCard").Msg("Finish")
		}
	}()
	return d.base.DeleteCard(ctx, i1)
}

// GetCardByID implements repository.CardRepo
func (d CardRepoWithLogs) GetCardByID(ctx context.Context, i1 int64) (c2 models.CardTable, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"i1":  i1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"c2":  c2,
				"err": err}).Err(err).Str("decorator", "CardRepoWithLogs").Str("method", "GetCardByID").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"c2":  c2,
				"err": err}).Str("decorator", "CardRepoWithLogs").Str("method", "GetCardByID").Msg("Finish")
		}
	}()
	return d.base.GetCardByID(ctx, i1)
}

// GetCardByName implements repository.CardRepo
func (d CardRepoWithLogs) GetCardByName(ctx context.Context, s1 string) (c2 models.CardTable, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"s1":  s1}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"c2":  c2,
				"err": err}).Err(err).Str("decorator", "CardRepoWithLogs").Str("method", "GetCardByName").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"c2":  c2,
				"err": err}).Str("decorator", "CardRepoWithLogs").Str("method", "GetCardByName").Msg("Finish")
		}
	}()
	return d.base.GetCardByName(ctx, s1)
}

// InsertCard implements repository.CardRepo
func (d CardRepoWithLogs) InsertCard(ctx context.Context, c2 models.CardTable) (i1 int64, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"c2":  c2}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Err(err).Str("decorator", "CardRepoWithLogs").Str("method", "InsertCard").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Str("decorator", "CardRepoWithLogs").Str("method", "InsertCard").Msg("Finish")
		}
	}()
	return d.base.InsertCard(ctx, c2)
}

// UpdateCard implements repository.CardRepo
func (d CardRepoWithLogs) UpdateCard(ctx context.Context, c2 models.CardTable) (i1 int64, err error) {

	nl := zerolog.Ctx(ctx)

	logger := nl.With().Fields(map[string]interface{}{
		"ctx": ctx,
		"c2":  c2}).Logger()

	defer func() {
		if err != nil {
			logger.Error().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Err(err).Str("decorator", "CardRepoWithLogs").Str("method", "UpdateCard").Msg("Error detected")
		} else {
			logger.Debug().Fields(map[string]interface{}{
				"i1":  i1,
				"err": err}).Str("decorator", "CardRepoWithLogs").Str("method", "UpdateCard").Msg("Finish")
		}
	}()
	return d.base.UpdateCard(ctx, c2)
}

// NewCardRepoWithLogs instruments an implementation of the repository.CardRepo with simple logging
func NewCardRepoWithLogs(base repository.CardRepo) repository.CardRepo {
	decorate := os.Getenv("DECORATE")
	if decorate == "true" || decorate == "1" {
		return CardRepoWithLogs{
			base: base,
		}
	}

	return base
}
