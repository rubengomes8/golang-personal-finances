package instrumentation

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

var (
	Registry *prometheus.Registry

	Logger *zerolog.Logger
)

func Init() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Logger()
	Logger = &logger
	zerolog.DefaultContextLogger = Logger

	Registry = prometheus.NewRegistry()
}

func RegistryHandler() http.Handler {
	return promhttp.HandlerFor(Registry, promhttp.HandlerOpts{
		Registry: Registry,
	})
}
