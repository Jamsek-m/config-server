package middlewares

import "net/http"

var PoweredByFilter = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("X-Powered-By", "ConfigServer v1.0")
		next.ServeHTTP(res, req)
	})
}
