package middleware

import (
    "net/http"
    "project/utils"
    "strings"
    "context"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing auth token", http.StatusUnauthorized)
            return
        }
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        userID, err := utils.ParseJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "userID", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
