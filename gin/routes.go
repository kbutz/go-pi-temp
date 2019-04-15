package gin

import (
	"github.com/kbutz/go-pi-temp/config"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// InitRoutes : Creates all of the routes for our application and returns a router
func InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*", //cfg.Origins,
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		// ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true, //Should be true for production. - is more secure because we validate headers as opposed to ember.
	}))

	if config.Debug {
		debugGroup := router.Group("/debug")
		setDebugRoutes(debugGroup)
	}

	v1 := router.Group("/v1")
	{
		setTemperatureReadingRoutes(v1)
	}

	return router
}

func setTemperatureReadingRoutes(g *gin.RouterGroup) {
	g.GET("/temperature", temperatureReadingHandler)
}

func setDebugRoutes(g *gin.RouterGroup) {
	g.GET("/test")
}
