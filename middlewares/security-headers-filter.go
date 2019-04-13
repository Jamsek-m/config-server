package middlewares

import "net/http"

var SecurityHeadersFilter = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("X-Content-Type-Options", "nosniff")
		res.Header().Add("X-XSS-Protection", "1; mode=block")
		res.Header().Add("X-Frame-Options", "DENY")
		res.Header().Add("Strict-Transport-Security", "max-age=63072000")
		res.Header().Add("Content-Security-Policy", "script-src 'self' 'unsafe-inline';style-src 'self' 'unsafe-inline';object-src 'none';default-src 'self';")
		next.ServeHTTP(res, req)
	})
}
