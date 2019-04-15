package server

import (
	"github.com/Jamsek-m/config-server/endpoints"
	"github.com/Jamsek-m/config-server/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func initMiddlewares(router *mux.Router) {
	router.Use(middlewares.PoweredByFilter)
	router.Use(middlewares.SecurityHeadersFilter)
}

func initRoutes(router *mux.Router) {
	userEndpoint := endpoints.UserEndpoint{}
	configEndpoint := endpoints.ConfigEndpoint{}
	roleEndpoint := endpoints.RoleEndpoint{}
	authEndpoint := endpoints.AuthEndpoint{}

	router.HandleFunc("/v1/auth/login", authEndpoint.Login).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/logout", authEndpoint.Logout).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/is-authorized", authEndpoint.IsAuthorized).Methods(http.MethodGet)
	router.HandleFunc("/v1/auth/user", authEndpoint.GetCurrentUser).Methods(http.MethodGet)

	router.HandleFunc("/v1/users",
		userEndpoint.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/{id}",
		userEndpoint.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/v1/users",
		userEndpoint.CreateUser).Methods(http.MethodPost)

	router.HandleFunc("/v1/roles", roleEndpoint.GetRoles).Methods(http.MethodGet)

	router.PathPrefix("/v1/keys").HandlerFunc(configEndpoint.GetConfigByKey).Methods(http.MethodGet)
	router.PathPrefix("/v1/keys").HandlerFunc(configEndpoint.CreateConfigEntry).Methods(http.MethodPut)

}
