package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CyberSleeper/backend-oprec-ristek/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010\n")
	c := cors.AllowAll()
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:9010", handler))
}
