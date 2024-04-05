package v1

import (
	"context"
	"net/http"

	api "github.com/vira-software/auth-server/internal/controller/http"
	"github.com/vira-software/auth-server/internal/log"
	services "github.com/vira-software/auth-server/internal/services"
)

// AuthMiddleware creates a middleware that verifies access token.
// It returns api.Middleware instance.
func AuthMiddleware(auth services.Auth, logger log.Logger) api.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := readAccessToken(r)
			if err != nil {
				api.ErrorJSON(w, err.Error(), http.StatusUnauthorized)
				return
			}
			fingerprint := readFingerprint(r)
			userID, roleTitles, err := auth.Verify(token, fingerprint)
			if err != nil {
				api.HandleError(err, func(e *services.Error) {
					api.ErrorJSON(w, e.Err.Error(), http.StatusUnauthorized)
				})
				return
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "roleTitles", roleTitles)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
