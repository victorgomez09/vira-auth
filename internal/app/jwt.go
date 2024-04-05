package app

import (
	"fmt"

	"github.com/vira-software/auth-server/internal/jwt"
)

// NewJWT reads private and public keys, creates JWT builder and JWT parser.
// It returns error if keys read failed or configuration is incorrect.
func NewJWT(cfg *Config) (jwt.Builder, jwt.Parser, error) {
	privateKey, err := jwt.ReadPrivateKey(cfg.Key.PrivatePath, cfg.AT.Alg)
	if err != nil {
		return nil, nil, fmt.Errorf("private key: %w", err)
	}

	publicKey, err := jwt.ReadPublicKey(cfg.Key.PublicPath, cfg.AT.Alg)
	if err != nil {
		return nil, nil, fmt.Errorf("public key: %w", err)
	}

	builder, err := jwt.NewBuilder(jwt.Params{Issuer: cfg.Name, Algorithm: cfg.AT.Alg, Key: privateKey})
	if err != nil {
		return nil, nil, fmt.Errorf("builder: %w", err)
	}

	parser, err := jwt.NewParser(jwt.Params{Issuer: cfg.Name, Algorithm: cfg.AT.Alg, Key: publicKey})
	if err != nil {
		return nil, nil, fmt.Errorf("parser: %w", err)
	}

	return builder, parser, nil
}
