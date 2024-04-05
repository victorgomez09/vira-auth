// Package usecase provides interfaces and structures to encapsulate business logic.
package usecase

import (
	"github.com/vira-software/auth-server/internal/models"
	"github.com/vira-software/auth-server/internal/uuid"
)

// User is interface implemented by types
// that can encapsulate user logic.
type User interface {
	// Create validates data and creates a new user.
	// It returns pointer to an models.User instance.
	Create(data models.User) (*models.User, error)

	// Get gets a user by ID.
	// It returns pointer to an models.User instance.
	Get(id uuid.UUID) (*models.User, error)

	// Verify verifies user's name and password
	// and is used in authentication process.
	// It returns user ID if name and password are correct.
	Verify(data models.User) (uuid.UUID, error)

	// UpdatePassword updates user's password by user ID
	// if user exists and currentPassword is correct.
	UpdatePassword(id uuid.UUID, currentPassword []byte, newPassword []byte) error

	// Delete deletes a user by ID
	// if user exists and currentPassword is correct.
	Delete(id uuid.UUID, currentPassword []byte) error
}

// Token is interface implemented by types
// that can encapsulate token logic.
type Token interface {
	// Create creates new access and refresh tokens using user's fingerprint.
	// It returns models.AccessToken instance
	// and pointer to an models.RefreshToken instance.
	Create(userID uuid.UUID, fingerprint []byte, session bool) (models.AccessToken, *models.RefreshToken, error)

	// Refresh verifies user's fingerprint and current refresh token by ID
	// and creates new access and refresh tokens.
	// It returns models.AccessToken instance
	// and pointer to an models.RefreshToken instance.
	Refresh(id uuid.UUID, fingerprint []byte) (models.AccessToken, *models.RefreshToken, error)

	// Get gets a refresh token by ID.
	// It returns pointer to an models.RefreshToken instance
	// if id is correct and token isn't expired.
	Get(id uuid.UUID) (*models.RefreshToken, error)

	// Delete verifies user's fingerprint and current refresh token by ID
	// and deletes a refresh token by ID.
	// It returns error if id is incorrect or token is expired.
	Delete(id uuid.UUID, fingerprint []byte) error

	// DeleteAll verifies user's fingerprint and current refresh token by ID
	// and deletes all user refresh tokens.
	// It returns error if id is incorrect or token is expired.
	DeleteAll(id uuid.UUID, fingerprint []byte) error
}

// Auth is interface implemented by types
// that can encapsulate authorization logic.
type Auth interface {
	// Verify verifies user's fingerprint, parses access token,
	// and retrieves user ID and roles from it.
	// It returns user ID and roles if token is correct and not expired.
	Verify(token models.AccessToken, fingerprint []byte) (uuid.UUID, []string, error)
}
