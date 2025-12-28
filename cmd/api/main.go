package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/darelife/calorieGo/internal/cache"
	"github.com/darelife/calorieGo/internal/config"
	"github.com/darelife/calorieGo/internal/external"
	"github.com/darelife/calorieGo/internal/handlers"
	"github.com/darelife/calorieGo/internal/middleware"
	"github.com/darelife/calorieGo/internal/server"
)

func main() {
	// -------- Logger --------
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// -------- Config --------
	cfg := config.Load()

	// -------- Redis --------
	redisCache := cache.New(cfg.Redis.Addr)
	if err := redisCache.Ping(context.Background()); err != nil {
		logger.Fatal("failed to connect to redis", zap.Error(err))
	}
	logger.Info("connected to redis")

	// -------- External API --------
	foodAPI := external.NewHTTPFoodAPI(cfg.External.BaseURL)

	// -------- Handlers --------
	foodHandler := handlers.NewFoodHandler(
		logger,
		redisCache,
		foodAPI,
	)

	// -------- Router --------
	r := chi.NewRouter()
	r.Use(middleware.Logging(logger))

	r.Post("/v1/food/barcode", foodHandler.GetByBarcode)

	// -------- HTTP Server --------
	srv := server.New(":"+cfg.Port, r, logger)

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server crashed", zap.Error(err))
		}
	}()

	// -------- Graceful Shutdown --------
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", zap.Error(err))
	}
}

