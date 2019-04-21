package middlewares

import (
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/response"
	"net/http"
	"strings"
)

/*
 CORS configuration model
*/
type corsConfigModel struct {
	origins           []string
	supportedMethods  []string
	_supportedMethods string
	supportedHeaders  []string
	_supportedHeaders string
	exposedHeaders    []string
	_exposedHeaders   string
}

func newCorsConfigModel(conf config.Config) corsConfigModel {
	c := corsConfigModel{}
	c.origins = strings.Split(conf.Server.Cors.AllowedOrigins, ",")
	c.supportedMethods = strings.Split(conf.Server.Cors.SupportedMethods, ",")
	c._supportedMethods = conf.Server.Cors.SupportedMethods
	c.supportedHeaders = strings.Split(conf.Server.Cors.SupportedHeaders, ",")
	c._supportedHeaders = conf.Server.Cors.SupportedHeaders
	c.exposedHeaders = strings.Split(conf.Server.Cors.ExposedHeaders, ",")
	c._exposedHeaders = conf.Server.Cors.ExposedHeaders
	return c
}

func (c corsConfigModel) getOrigin(origin string) string {
	for _, o := range c.origins {
		if origin == o {
			return origin
		} else if o == "*" {
			return origin
		}
	}
	return ""
}

func (c corsConfigModel) getSupportedMethod(method string) string {
	if method == "" {
		method = http.MethodGet
	}
	for _, m := range c.supportedMethods {
		if method == m {
			return c._supportedMethods
		} else if m == "*" {
			return c._supportedMethods
		}
	}
	return ""
}

func (c corsConfigModel) getSupportedHeaders(header string) string {
	for _, h := range c.supportedHeaders {
		if header == h {
			return c._supportedHeaders
		} else if h == "*" {
			return c._supportedHeaders
		}
	}
	return ""
}

func (c corsConfigModel) getExposedHeaders(header string) string {
	for _, h := range c.exposedHeaders {
		if header == h {
			return c._exposedHeaders
		} else if h == "*" {
			return c._exposedHeaders
		}
	}
	return ""
}

var corsConfig corsConfigModel

/*
 Initialization function for CORS configuration
*/
func InitializeCorsFilter() {
	corsConfig = newCorsConfigModel(config.GetConfiguration())
}

/*
 CORS filter
*/
var CorsFilter = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		allOk := handleCors(req, res)
		if allOk {
			next.ServeHTTP(res, req)
		} else {
			return
		}
	})
}

func handleCors(req *http.Request, res http.ResponseWriter) bool {
	httpMethod := corsConfig.getSupportedMethod(req.Method)
	if httpMethod == "" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	origin := req.Header.Get(response.ORIGIN)
	if origin != "" {
		allowedOrigin := corsConfig.getOrigin(origin)
		if allowedOrigin != "" {
			res.Header().Add(response.ACCESS_CONTROL_ALLOW_ORIGIN, allowedOrigin)
		}
		// TODO: handle headers if necessary
	}
	return true
}
