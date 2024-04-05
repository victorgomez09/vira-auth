// Package repo provides interfaces and structures to interact with database.
package repo

import (
	"context"

	"github.com/vira-software/auth-server/internal/models"
	"github.com/vira-software/auth-server/internal/uuid"
)

// Role is interface implemented by types
// that can interact with role entity.
type Role interface {
	// Create creates a new role.
	// It returns pointer to an entity.Role instance.
	Create(ctx context.Context, data models.Role) (*models.Role, error)

	// GetByID gets a role by ID.
	// It returns pointer to an entity.Role instance.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Role, error)

	// GetByUser gets roles by user ID.
	// It returns slice of entity.Role instances.
	GetByUser(ctx context.Context, userID uuid.UUID) ([]models.Role, error)

	// DeleteByID deletes a role by ID.
	DeleteByID(ctx context.Context, id uuid.UUID) error

	// DeleteByUser deletes user-related roles by user ID.
	DeleteByUser(ctx context.Context, userID uuid.UUID) error
}

// User is interface implemented by types
// that can interact with user entity.
type User interface {
	// Create creates a new user.
	// It returns pointer to an entity.User instance.
	Create(ctx context.Context, data models.User) (*models.User, error)

	// GetByID gets a user by ID.
	// It returns pointer to an entity.User instance.
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)

	// GetByName gets a user by unique name.
	// It returns pointer to an entity.User instance.
	GetByName(ctx context.Context, name string) (*models.User, error)

	// UpdatePassword updates user's password by user ID.
	UpdatePassword(ctx context.Context, id uuid.UUID, password []byte) error

	// DeleteByID deletes a user by ID.
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

// Token is interface implemented by types
// that can interact with token entity.
type Token interface {
	// Create creates a new refresh token.
	// It returns pointer to an entity.RefreshToken instance.
	Create(ctx context.Context, data models.RefreshToken) (*models.RefreshToken, error)

	// GetByID gets a refresh token by ID.
	// It returns pointer to an entity.RefreshToken instance.
	GetByID(ctx context.Context, id uuid.UUID) (*models.RefreshToken, error)

	// GetByUser gets refresh tokens by user ID.
	// It returns slice of entity.RefreshToken instances.
	GetByUser(ctx context.Context, userID uuid.UUID) ([]models.RefreshToken, error)

	// DeleteByID deletes a refresh token by ID.
	DeleteByID(ctx context.Context, id uuid.UUID) error

	// DeleteByUser deletes user-related refresh tokens by user ID.
	DeleteByUser(ctx context.Context, userID uuid.UUID) error
}
