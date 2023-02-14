// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/red_template.go.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package user

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

type UserRepoWithRED struct {
	base         repository.UserRepo
	histogramVec *prometheus.HistogramVec
}

// GetUserByUsername implements repository.UserRepo
func (d UserRepoWithRED) GetUserByUsername(ctx context.Context, s1 string) (u1 models.UserTable, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "GetUserByUsername",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.GetUserByUsername(ctx, s1)
}

// InsertUser implements repository.UserRepo
func (d UserRepoWithRED) InsertUser(ctx context.Context, u1 models.UserTable) (i1 int64, err error) {
	since := time.Now()
	defer func() {
		status := "ok"
		if err != nil {
			status = "error"
		}

		labels := prometheus.Labels{
			"status": status,
			"method": "InsertUser",
		}

		observer, err := d.histogramVec.GetMetricWith(labels)
		if err != nil {
			fmt.Printf("Metric: Error to get metric with labels %v\n", labels)
		}

		observer.Observe(float64(time.Since(since).Milliseconds()))
	}()
	return d.base.InsertUser(ctx, u1)
}

// NewUserRepoWithRED returns an instance of the repository.UserRepo decorated with red histogram metric
func NewUserRepoWithRED(base repository.UserRepo, constLabels prometheus.Labels) (decorator repository.UserRepo, err error) {
	decorate := os.Getenv("DECORATE")
	if !(decorate == "true" || decorate == "1") {
		return base, nil
	}

	subSystem := "user_repo"

	metricConfig := prometheus.HistogramOpts{
		Namespace:   strings.TrimSpace("system"),
		Subsystem:   subSystem,
		Name:        fmt.Sprintf("%s_red", subSystem),
		Help:        "UserRepo RED histogram (rate, errors and duration).",
		ConstLabels: constLabels,
		Buckets:     prometheus.ExponentialBuckets(100, 2, 5),
	}

	red := UserRepoWithRED{
		base:         base,
		histogramVec: prometheus.NewHistogramVec(metricConfig, []string{"status", "method"}),
	}

	err = instrumentation.Registry.Register(red.histogramVec)
	if err != nil {
		return nil, err
	}

	return red, nil
}
