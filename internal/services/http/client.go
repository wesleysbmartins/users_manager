package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"users_manager/internal/domain/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HttpClient() {
	router := mux.NewRouter()
	router.Use(ContentTypeMiddleware)
	routes.Routes(router)
	http.Handle("/", router)

	port := os.Getenv("PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{allowedOrigins}))(router)

	fmt.Println(fmt.Sprintf("Server listenning on Port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler))
}
