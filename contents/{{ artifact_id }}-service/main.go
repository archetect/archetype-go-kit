package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"{{artifact_id}}-service/{{artifact_id}}"
)

func main() {

	var (
		httpAddr = flag.String("http", ":{{default_port}}", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// setting up logger with timestamp and a prefix as application name for easier filtering
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "{{artifact_id}}", log.DefaultTimestamp)

	// Creating and Configuring Service with logging and monitoring
	fieldKeys := []string{"method", "error"} //Need this for prometheus
	//Creating new instance of {{artifact_id}} service
	svc := {{artifact_id}}.NewService()
	svc = {{artifact_id}}.NewLoggingService(log.With(logger, "component", "{{artifact_id}}"), svc)
	svc = {{artifact_id}}.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "{{artifact_id}}",
			Subsystem: "{{artifact_id}}_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "{{artifact_id}}",
			Subsystem: "{{artifact_id}}_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "{{artifact_id}}",
			Subsystem: "{{artifact_id}}_service",
			Name:      "count_result",
			Help:      "The result of each status method.",
		}, fieldKeys),
		svc,
	)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := {{artifact_id}}.Endpoints{
		StatusEndpoint:       {{artifact_id}}.MakeStatusEndpoint(svc),
	}

	// HTTP transport
	go func() {
		logger.Log("{{artifact_id}}-service is listening on port:", *httpAddr)
		handler := {{artifact_id}}.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	logger.Log(<-errChan)
}

