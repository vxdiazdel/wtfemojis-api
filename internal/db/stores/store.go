package stores

import (
	"context"

	"github.com/vxdiazdel/wtfemojis-api/models"
)

type IStore interface {
	GetEmojis(ctx context.Context, name string) ([]models.Emoji, error)
}
