// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/red_template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package expense

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

type ExpenseRepoWithRED struct {
	base         repository.ExpenseRepo
	histogramVec *prometheus.HistogramVec
}

// DeleteExpense implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) DeleteExpense(ctx context.Context, i1 int64) (err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "DeleteExpense",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.DeleteExpense(ctx, i1)
}

// GetExpenseByID implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) GetExpenseByID(ctx context.Context, i1 int64) (e1 models.ExpenseView, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetExpenseByID",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetExpenseByID(ctx, i1)
}

// GetExpensesByCard implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) GetExpensesByCard(ctx context.Context, s1 string) (ea1 []models.ExpenseView, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetExpensesByCard",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetExpensesByCard(ctx, s1)
}

// GetExpensesByCategory implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) GetExpensesByCategory(ctx context.Context, s1 string) (ea1 []models.ExpenseView, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetExpensesByCategory",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetExpensesByCategory(ctx, s1)
}

// GetExpensesByDates implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) GetExpensesByDates(ctx context.Context, t1 time.Time, t2 time.Time) (ea1 []models.ExpenseView, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetExpensesByDates",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetExpensesByDates(ctx, t1, t2)
}

// GetExpensesBySubCategory implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) GetExpensesBySubCategory(ctx context.Context, s1 string) (ea1 []models.ExpenseView, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetExpensesBySubCategory",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetExpensesBySubCategory(ctx, s1)
}

// InsertExpense implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) InsertExpense(ctx context.Context, e1 models.ExpenseTable) (i1 int64, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "InsertExpense",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.InsertExpense(ctx, e1)
}

// UpdateExpense implements repository.ExpenseRepo
func (d ExpenseRepoWithRED) UpdateExpense(ctx context.Context, e1 models.ExpenseTable) (i1 int64, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "UpdateExpense",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.UpdateExpense(ctx, e1)
}

// NewExpenseRepoWithRED returns an instance of the repository.ExpenseRepo decorated with red histogram metric
func NewExpenseRepoWithRED(base repository.ExpenseRepo, constLabels prometheus.Labels) (decorator repository.ExpenseRepo, err error) {
	decorate := os.Getenv("DECORATE")
	if !(decorate == "true" || decorate == "1") {
		return base, nil
	}

	subSystem := "expense_repo"

	metricConfig := prometheus.HistogramOpts{
		Namespace:   strings.TrimSpace("system"),
		Subsystem:   subSystem,
		Name:        fmt.Sprintf("%s_red", subSystem),
		Help:        "ExpenseRepo RED histogram (rate, errors and duration).",
		ConstLabels: constLabels,
		Buckets:     prometheus.ExponentialBuckets(100, 2, 5),
	}

	red := ExpenseRepoWithRED{
		base:         base,
		histogramVec: prometheus.NewHistogramVec(metricConfig, []string{"status", "method"}),
	}

	err = instrumentation.Registry.Register(red.histogramVec)
	if err != nil {
		return nil, err
	}

	return red, nil
}
