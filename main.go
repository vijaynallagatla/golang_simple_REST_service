package main

import (
	"net/http"
	"os"

	applicationRoute "golang_simle_REST_service/router"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func logMessageEveryTenSecond() {
	log.WithFields(log.Fields{
		"tag": "INFO",
	}).Info("Sample goLang REST service.")
}
func main() {
	go logMessageEveryTenSecond()
	router := httprouter.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://127.0.0.1:8081"},
		AllowCredentials: false,
		Debug:            false,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	applicationRoute.ApplicationRoutes(router)
	log.Printf("Server started. \n Listening at 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
