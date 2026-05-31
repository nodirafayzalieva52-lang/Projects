package middleware

import (
	"net/http"
	"project/logger"

	"go.uber.org/zap"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logger.L.Info("Logging middleware",zap.String("method",r.Method), zap.String("path",r.URL.Path))

		next.ServeHTTP(w, r)
	})
}
