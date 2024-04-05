// Package model provides model structures for use cases, repositories and controllers.
package models

import "github.com/vira-software/auth-server/internal/uuid"

// Role model.
type Role struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
