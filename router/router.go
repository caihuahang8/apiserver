package router

import (
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	// 404 handler
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
