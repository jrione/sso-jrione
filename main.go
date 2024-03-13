package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	middleware "github.com/jrione/sso-jrione/middleware"
	"github.com/jrione/sso-jrione/route"
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

	if env.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	route.SetupRoute(env, timeout, dbclient, r)
	r.Run(env.Server.Listen + ":" + env.Server.Port)
}
