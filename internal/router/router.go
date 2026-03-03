package router

import (
	"github.com/gin-gonic/gin"
	"github.com/My-second-free-organization/platform-api/internal/config"
	"github.com/My-second-free-organization/platform-api/internal/handler"
	"github.com/My-second-free-organization/platform-api/internal/middleware"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.CORS(), middleware.RequestID())
	r.GET("/health", handler.HealthCheck)
	r.GET("/ready", handler.ReadinessCheck)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth(cfg.JWTSecret))
	wf := v1.Group("/workflows")
	{ wh := handler.NewWorkflowHandler(cfg); wf.POST("", wh.Create); wf.GET("", wh.List); wf.GET("/:id", wh.Get); wf.DELETE("/:id", wh.Delete) }
	tasks := v1.Group("/tasks")
	{ th := handler.NewTaskHandler(cfg); tasks.GET("/:id", th.Get); tasks.POST("/:id/complete", th.Complete) }
	return r
}
