package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	cli "github.com/jawher/mow.cli"
	metrics "github.com/rcrowley/go-metrics"

	api "github.com/Financial-Times/api-endpoint"
	fthealth "github.com/Financial-Times/go-fthealth/v1_1"
	logger "github.com/Financial-Times/go-logger/v2"
	"github.com/Financial-Times/http-handlers-go/v2/httphandlers"
	status "github.com/Financial-Times/service-status-go/httphandlers"
)

const (
	// TODO: set meaningful appDescription
	appDescription = ""
	appDefaultName = "cm-go-service"
	// TODO: how long we would like to wait for response?
	httpServerReadTimeout  = 10 * time.Second
	httpServerWriteTimeout = 15 * time.Second
	httpServerIdleTimeout  = 20 * time.Second
	httpHandlersTimeout    = 14 * time.Second
)

func main() {
	app := cli.App(appDefaultName, appDescription)

	appSystemCode := app.String(cli.StringOpt{
		Name:   "app-system-code",
		Value:  "cm-go-service", // TODO: update to match the Biz-Ops system code
		Desc:   "system Code of the application",
		EnvVar: "APP_SYSTEM_CODE",
	})

	appName := app.String(cli.StringOpt{
		Name:   "app-name",
		Value:  appDefaultName,
		Desc:   "application name",
		EnvVar: "APP_NAME",
	})

	port := app.String(cli.StringOpt{
		Name:   "port",
		Value:  "8080",
		Desc:   "port to listen on",
		EnvVar: "APP_PORT",
	})

	logLevel := app.String(cli.StringOpt{
		Name:   "log-level",
		Value:  "INFO",
		Desc:   "logging level (DEBUG, INFO, WARN, ERROR)",
		EnvVar: "LOG_LEVEL",
	})

	apiYML := app.String(cli.StringOpt{
		Name:   "api-yml",
		Value:  "./api/api.yml",
		Desc:   "Location of the OpenAPI YML file.",
		EnvVar: "API_YML",
	})

	log := logger.NewUPPLogger(*appName, *logLevel)

	app.Action = func() {
		log.Infof("Starting with system code: %s, app name: %s, port: %s", *appSystemCode, *appName, *port)

		hc := HealthConfig{
			appSystemCode:  *appSystemCode,
			appName:        *appName,
			appDescription: appDescription,
		}

		healthService := NewHealthService(hc)

		apiEndpoint, err := newAPIEndpoint(*apiYML)
		if err != nil {
			log.WithError(err).WithField("file", *apiYML).
				Warn("Failed to serve the API Endpoint for this service. Please validate the file exists, and that it fits the OpenAPI specification.")
		}

		router := registerEndpoints(healthService, apiEndpoint, log)

		server := newHTTPServer(*port, router)
		go startHTTPServer(server, log)

		waitForSignal()
		stopHTTPServer(server, log)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("App could not start: %v", err)
		return
	}
}

func newAPIEndpoint(apiFileName string) (api.Endpoint, error) {
	if apiFileName == "" {
		return nil, nil
	}

	apiEndpoint, err := api.NewAPIEndpointForFile(apiFileName)
	if err != nil {
		return nil, err
	}

	return apiEndpoint, nil
}

func registerEndpoints(healthService *HealthService, apiEndpoint api.Endpoint, log *logger.UPPLogger) http.Handler {
	serveMux := http.NewServeMux()

	// register supervisory endpoint that does not require logging and metrics collection
	serveMux.HandleFunc("/__health", fthealth.Handler(healthService.Health()))
	serveMux.HandleFunc(status.GTGPath, status.NewGoodToGoHandler(healthService.GTG))
	serveMux.HandleFunc(status.BuildInfoPath, status.BuildInfoHandler)

	if apiEndpoint != nil {
		serveMux.HandleFunc(api.DefaultPath, apiEndpoint.ServeHTTP)
	}

	// add services router and register endpoints specific to this service only
	servicesRouter := mux.NewRouter()
	//TODO: add real handlers
	servicesRouter.HandleFunc("/test", TestHandler).Methods("GET")

	// wrap the handler with certain middlewares providing logging of the requests,
	// sending metrics and handler time out on certain time interval
	var wrappedServicesRouter http.Handler = servicesRouter
	wrappedServicesRouter = httphandlers.TransactionAwareRequestLoggingHandler(log, wrappedServicesRouter)
	wrappedServicesRouter = httphandlers.HTTPMetricsHandler(metrics.DefaultRegistry, wrappedServicesRouter)
	wrappedServicesRouter = http.TimeoutHandler(wrappedServicesRouter, httpHandlersTimeout, "")

	serveMux.Handle("/", wrappedServicesRouter)

	return serveMux
}

func newHTTPServer(port string, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
		IdleTimeout:  httpServerIdleTimeout,
	}
}

func startHTTPServer(server *http.Server, log *logger.UPPLogger) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}

func stopHTTPServer(server *http.Server, log *logger.UPPLogger) {
	log.Info("HTTP server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown the server: %v", err)
	}
}

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
