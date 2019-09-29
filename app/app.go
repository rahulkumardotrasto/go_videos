package app

import (
	"math/rand"
	"time"

	gzip "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//App has router and db instances
type App struct{}

//Init initializes the app with predefined configuration
func (app *App) Init() {
	rand.Seed(time.Now().UnixNano())
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.Use(gin.ErrorLogger())
	routes.Use(gin.Recovery())

	routes.StaticFile("/favicon.ico", "")
	routes.Run(":8000")
}
