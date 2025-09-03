package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func OTLPMiddleware(env *config.Config, ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(env.Jaeger.Host+":"+env.Jaeger.Port),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(env.Server.AppName),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func AuthMiddleware(env *config.Config) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		apiKey := gctx.GetHeader("X-API-Key")
		if apiKey != env.Server.XApiKey {
			gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token X-Api-Key not provided",
			})
			gctx.Abort()
			return
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		gctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		gctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		gctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if gctx.Request.Method == "OPTIONS" {
			gctx.AbortWithStatus(204)
			return
		}

		gctx.Next()
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		sess := sessions.Default(gctx)
		isLoggedIn := sess.Get("IsLoggedIn")
		fmt.Println("sess", isLoggedIn)
		if isLoggedIn != true {
			gctx.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			gctx.Abort()
		} else {
			gctx.Next()
		}
	}
}

func JWTMiddleware(tokenSecret string) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req := gctx.Request.Header.Get("Authorization")
		t := strings.Split(req, " ")
		if len(req) != 0 {
			authToken := t[1]
			ok, err := config.IsAuthorized(authToken, tokenSecret)
			if err != nil {
				gctx.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Internal Status Error",
					"Cause": err.Error(),
				})
				gctx.Abort()
				return
			}
			if !ok {
				gctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Bearer token is missing",
				})
				gctx.Abort()
				return
			}
			return
		} else {
			gctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			gctx.Abort()
			return
		}
	}
}
