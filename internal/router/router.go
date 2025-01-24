package router

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/wtfemojis-api/handlers"
	"github.com/vxdiazdel/wtfemojis-api/internal/db/stores"
)

func NewRouter(
	ctx context.Context,
	store stores.IStore,
) *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://wtfemojis.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	h := handlers.NewHandlerContext(
		ctx,
		store,
	)

	api := r.Group("/api/v1")

	emojis := api.Group("/emojis")
	emojis.GET("", h.GetEmojis)

	return r
}
