package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/neekonsu/personalwebsite/config"
	"github.com/neekonsu/personalwebsite/handlers"
)

func main() {
	cfg := config.Load()

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root route and other pages
	http.HandleFunc("/", handlers.ServeHTML)

	// Start the server
	fmt.Printf("Server is running on http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
