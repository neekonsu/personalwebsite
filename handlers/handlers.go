package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func getStaticDir() string {
	if dir := os.Getenv("STATIC_DIR"); dir != "" {
		return dir
	}
	return "static"
}

// ServeHTML handles serving HTML pages
func ServeHTML(w http.ResponseWriter, r *http.Request) {
	staticDir := getStaticDir()

	if r.URL.Path == "/" {
		// Serve the main HTML file
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
		return
	}

	// For other routes, serve the corresponding HTML file
	path := filepath.Join(staticDir, "pages", r.URL.Path+".html")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the file doesn't exist, serve a 404 page
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}
