package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/controllers"
	"server/middleware"
)

func RunServer() {
	server := http.Server{
		Handler: newRoute(),
		Addr:    ":8000",
	}

	log.Printf("Starting up server on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func newRoute() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/registration", controllers.Register).Methods("POST")

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware)
	apiRouter.HandleFunc("/files", controllers.Upload).Methods("POST")
	apiRouter.HandleFunc("/files/{fileName}", controllers.Download).Methods("GET")
	apiRouter.HandleFunc("/files", controllers.Index).Methods("GET")

	return router
}
