package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/My-second-free-organization/platform-api/internal/config"
	"github.com/My-second-free-organization/platform-api/internal/router"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, _ := config.Load()
	r := router.New(cfg)
	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.ServerPort), Handler: r, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second}
	go func() { log.Info().Int("port", cfg.ServerPort).Msg("Listening"); srv.ListenAndServe() }()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
