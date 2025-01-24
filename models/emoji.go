package models

import "github.com/google/uuid"

type Emoji struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}
