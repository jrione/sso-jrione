package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func OTelMetrics(env *config.Config) gin.HandlerFunc {

	var (
		meter           = otel.Meter(env.Server.AppName)
		httpRequests, _ = meter.Int64Counter("http_requests_total")
		httpLatency, _  = meter.Float64Histogram("http_request_duration_seconds")
	)

	return func(gctx *gin.Context) {
		start := time.Now()

		gctx.Next()

		latency := time.Since(start).Seconds()
		status := gctx.Writer.Status()

		labels := []attribute.KeyValue{
			attribute.String("http.method", gctx.Request.Method),
			attribute.String("http.route", gctx.FullPath()),
			attribute.Int("http.status_code", status),
		}

		httpRequests.Add(gctx.Request.Context(), 1, metric.WithAttributes(labels...))
		httpLatency.Record(gctx.Request.Context(), latency, metric.WithAttributes(labels...))
	}
}
