package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vxdiazdel/wtfemojis-api/internal/db"
	"github.com/vxdiazdel/wtfemojis-api/internal/db/stores"
	"github.com/vxdiazdel/wtfemojis-api/internal/router"
)

func init() {
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	ctx := context.Background()

	// load .env in development
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(fmt.Errorf("load .env: %w", err))
		}
	}

	// logger
	lg := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(lg)

	// database
	dbConn := db.NewPostgresConn(ctx, os.Getenv("DB_URL"))
	defer dbConn.Close(ctx)

	// stores
	store := stores.NewPostgresStore(ctx, dbConn)

	// router
	r := router.NewRouter(ctx, store)
	if err := r.Run(); err != nil {
		panic(fmt.Errorf("router run: %w", err))
	}
}
