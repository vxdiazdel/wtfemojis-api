package handlers

import (
	"context"

	"github.com/vxdiazdel/wtfemojis-api/internal/db/stores"
)

type HandlerContext struct {
	ctx   context.Context
	store stores.IStore
}

func NewHandlerContext(
	ctx context.Context,
	store stores.IStore,
) *HandlerContext {
	return &HandlerContext{
		ctx:   ctx,
		store: store,
	}
}

func (h *HandlerContext) Ctx() context.Context {
	return h.ctx
}

func (h *HandlerContext) Store() stores.IStore {
	return h.store
}
