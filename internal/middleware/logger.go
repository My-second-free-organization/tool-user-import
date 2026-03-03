package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now(); c.Next()
		log.Info().Str("method", c.Request.Method).Str("path", c.Request.URL.Path).Int("status", c.Writer.Status()).Dur("latency", time.Since(start)).Msg("req")
	}
}
func RequestID() gin.HandlerFunc { return func(c *gin.Context) { c.Set("request_id", "req-"+time.Now().Format("20060102150405")); c.Next() } }
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*"); c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" { c.AbortWithStatus(204); return }; c.Next()
	}
}
