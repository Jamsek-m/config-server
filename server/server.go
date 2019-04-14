package server

import (
	"fmt"
	"github.com/Jamsek-m/config-server/config"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func StartServer() {
	router := mux.NewRouter()
	initMiddlewares(router)
	initRoutes(router)

	port := config.GetConfiguration().Server.Port

	if port == 0 {
		port = 8000
	}

	fmt.Printf("Starting Config Server at port %d\n", port)
	serverAddr := fmt.Sprintf(":%d", port)
	serverError := http.ListenAndServe(serverAddr, router)
	if serverError != nil {
		fmt.Println(serverError)
		os.Exit(3)
	}

}
