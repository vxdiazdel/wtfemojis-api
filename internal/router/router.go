package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/wtfemojis-api/handlers"
	"github.com/vxdiazdel/wtfemojis-api/internal/db/stores"
)

func NewRouter(
	ctx context.Context,
	store stores.IStore,
) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandlerContext(
		ctx,
		store,
	)

	api := r.Group("/api/v1")

	emojis := api.Group("/emojis")
	emojis.GET("", h.GetEmojis)

	return r
}
