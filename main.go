package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	middleware "github.com/jrione/sso-jrione/middleware"
	"github.com/jrione/sso-jrione/route"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	app := config.Bootstrap()
	env := app.Env
	timeout := time.Duration(env.Server.Timeout) * time.Second

	dbclient := app.DBClient
	defer func() {
		if err := dbclient.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := context.Background()
	tp, err := middleware.OTLPMiddleware(env, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(nil); err != nil {
			log.Fatal(err)
		}
	}()

	if env.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(otelgin.Middleware(env.Server.AppName))
	route.SetupRoute(env, timeout, dbclient, r)
	if err := r.Run(env.Server.Listen + ":" + env.Server.Port); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
