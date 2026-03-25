package boot

import (
	"context"
	"net/http"

	docs "github.com/kweaver-ai/agent-observability/docs/swagger"
	"github.com/kweaver-ai/agent-observability/src/conf"
	"github.com/kweaver-ai/agent-observability/src/domain/service/tracesvc"
	"github.com/kweaver-ai/agent-observability/src/drivenadapter/httpaccess/opensearchtraceaccess"
	"github.com/kweaver-ai/agent-observability/src/driveradapter/api/httphandler"
	"github.com/kweaver-ai/agent-observability/src/infra/opensearch"
	"github.com/kweaver-ai/agent-observability/src/infra/server/httpserver"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type App struct {
	server *httpserver.Server
}

func NewApp() *App {
	httpServerConfig := conf.NewHTTPServerConfig()
	openSearchConfig := conf.NewOpenSearchConfig()
	docs.SwaggerInfo.BasePath = "/"

	openSearchClient := opensearch.New(
		openSearchConfig.Endpoint,
		opensearch.AuthConfig{
			Enabled:  openSearchConfig.Auth.Enabled,
			Username: openSearchConfig.Auth.Username,
			Password: openSearchConfig.Auth.Password,
		},
		openSearchConfig.Timeout,
	)
	traceDetailClient := opensearchtraceaccess.New(openSearchClient, openSearchConfig.TraceIndex)
	traceQueryService := tracesvc.New(traceDetailClient)
	traceHandler := httphandler.NewTraceHandler(traceQueryService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/traces/_search", traceHandler.SearchTraces)
	mux.HandleFunc("/api/v1/traces/by-conversation", traceHandler.SearchTracesByConversationID)
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	return &App{
		server: httpserver.New(httpServerConfig.Address, mux),
	}
}

func (a *App) Start() error {
	return a.server.Start()
}

func (a *App) Shutdown(ctx context.Context) error {
	if a.server != nil {
		return a.server.Shutdown(ctx)
	}

	return nil
}
