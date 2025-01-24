package stores

import (
	"context"
	"fmt"

	"github.com/vxdiazdel/wtfemojis-api/models"
)

func (s *PostgresStore) GetEmojis(
	ctx context.Context,
	name string,
) ([]models.Emoji, error) {
	var filter string
	if name != "" {
		filter = fmt.Sprintf("WHERE name LIKE '%%%s%%'", name)
	}

	q := fmt.Sprintf(`
		SELECT
			id,
			name,
			value
		FROM
			emojis
		%s
	`, filter)
	rows, err := s.DB().Query(s.Ctx(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emojis []models.Emoji
	for rows.Next() {
		var e models.Emoji
		if err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Value,
		); err != nil {
			return nil, err
		}

		emojis = append(emojis, e)
	}

	return emojis, nil
}
