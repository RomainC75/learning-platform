package middlewares

import (
	"context"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"net/http"
	"strings"
)

var (
	UserIdKey    = "user_id"
	UserEmailKey = "user_email"
	UserRoleKey  = "user_role"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		auth_header := r.Header.Get("Authorization")

		w.Header()
		if !strings.HasPrefix(auth_header, "Bearer") {
			http.Error(w, "token missing", http.StatusBadRequest)
			return
		}
		token = strings.Split(auth_header, " ")[1]

		claim, err := user_mngmt_infra.NewJWT().GetClaimsFromToken(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusBadRequest)
			return
		}

		bgCtx := context.Background()
		wvCtx := context.WithValue(bgCtx, UserIdKey, claim.UserId)
		wvCtx = context.WithValue(wvCtx, UserEmailKey, claim.UserEmail)
		wvCtx = context.WithValue(wvCtx, UserRoleKey, claim.UserRole)

		next.ServeHTTP(w, r.WithContext(wvCtx))
	})
}
