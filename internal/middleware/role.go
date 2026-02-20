package middleware

import (
	"context"
	"net/http"
	"strconv"

	"student-grade-management-system/internal/models"
)

type contextKey string

const (
	roleKey   contextKey = "role"
	userIDKey contextKey = "userID"
)

func RoleMiddleware(next http.Handler, allowedRoles ...models.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roleHeader := r.Header.Get("X-Role")
		userIDHeader := r.Header.Get("X-User-ID")

		if roleHeader == "" || userIDHeader == "" {
			http.Error(w, "missing X-Role or X-User-ID header", http.StatusBadRequest)
			return
		}

		role := models.Role(roleHeader)
		if !isValidRole(role) {
			http.Error(w, "invalid role", http.StatusForbidden)
			return
		}

		userID, err := strconv.Atoi(userIDHeader)
		if err != nil || userID <= 0 {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}

		if !isAllowedRole(role, allowedRoles) {
			http.Error(w, "access denied", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), roleKey, role)
		ctx = context.WithValue(ctx, userIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isValidRole(role models.Role) bool {
	return role == models.RoleAdmin ||
		role == models.RoleTeacher ||
		role == models.RoleStudent
}

func isAllowedRole(role models.Role, allowed []models.Role) bool {
	for _, r := range allowed {
		if role == r {
			return true
		}
	}
	return false
}

/* -------- Context Accessors -------- */

func GetRole(ctx context.Context) models.Role {
	if role, ok := ctx.Value(roleKey).(models.Role); ok {
		return role
	}
	return ""
}

func GetUserID(ctx context.Context) int {
	if id, ok := ctx.Value(userIDKey).(int); ok {
		return id
	}
	return 0
}
