package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
)

func NewTestRoute(env *config.Config, timeout time.Duration, gr *gin.RouterGroup) {
	gr.GET("/test", func(g *gin.Context) {
		g.JSON(200, gin.H{
			"njir": env.Server.Mode,
			"bruh": "bruh"
		})
	})
}
