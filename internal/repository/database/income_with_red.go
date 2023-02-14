package database

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rubengomes8/golang-personal-finances/internal/instrumentation"
	"github.com/rubengomes8/golang-personal-finances/internal/repository"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"
)

const (
	incomesSubsystem = "incomes"
)

// IncomesWithRED is the Incomes decorator adding the RED
type IncomesWithRED struct {
	repo         repository.IncomeRepo
	histogramVec *prometheus.HistogramVec
}

func (i IncomesWithRED) InsertIncome(ctx context.Context, income models.IncomeTable) (incomeID int64, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "InsertIncome",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.InsertIncome(ctx, income)
}
func (i IncomesWithRED) UpdateIncome(ctx context.Context, income models.IncomeTable) (incomeID int64, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "UpdateIncome",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.UpdateIncome(ctx, income)
}

func (i IncomesWithRED) GetIncomeByID(ctx context.Context, id int64) (income models.IncomeView, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "GetIncomeByID",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.GetIncomeByID(ctx, id)
}

func (i IncomesWithRED) GetIncomesByDates(ctx context.Context, minDate time.Time, maxDate time.Time) (income []models.IncomeView, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "GetIncomesByDates",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.GetIncomesByDates(ctx, minDate, maxDate)
}

func (i IncomesWithRED) GetIncomesByCategory(ctx context.Context, category string) (income []models.IncomeView, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "GetIncomesByCategory",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.GetIncomesByCategory(ctx, category)
}

func (i IncomesWithRED) GetIncomesByCard(ctx context.Context, card string) (income []models.IncomeView, err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "GetIncomesByCard",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.GetIncomesByCard(ctx, card)
}

func (i IncomesWithRED) DeleteIncome(ctx context.Context, id int64) (err error) {
	startingTime := time.Now().UTC()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		promLabels := prometheus.Labels{
			"status": status,
			"method": "DeleteIncome",
		}

		observer, err := i.histogramVec.GetMetricWith(promLabels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", promLabels)
		}

		observer.Observe(float64(time.Since(startingTime).Milliseconds()))
	}()

	return i.repo.DeleteIncome(ctx, id)
}

// NewIncomesWithRED returns an instance of the repository.IncomeRepo decorated with red histogram metric
func NewIncomesWithRED(base repository.IncomeRepo, constLabels prometheus.Labels) (repository.IncomeRepo, error) {
	decorate := os.Getenv("DECORATE")
	if !(decorate == "true" || decorate == "1") {
		return base, nil
	}

	metricConfig := prometheus.HistogramOpts{
		Namespace:   strings.TrimSpace("personal_finances"),
		Subsystem:   strings.TrimSpace(incomesSubsystem),
		Name:        strings.TrimSpace("database"),
		Help:        "Repository RED histogram (rate, errors and duration)",
		ConstLabels: constLabels,
		Buckets:     prometheus.ExponentialBuckets(100, 2, 5),
	}

	red := IncomesWithRED{
		repo:         base,
		histogramVec: prometheus.NewHistogramVec(metricConfig, []string{"status", "method"}),
	}

	err := instrumentation.Registry.Register(red.histogramVec)
	if err != nil {
		return nil, err
	}

	return red, nil
}
