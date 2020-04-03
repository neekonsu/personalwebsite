package main

import (
	"flag"
	"net/http"

	"github.com/neekonsu/personalwebsite"
)

func main() {
	html := flag.String("-html-path", "./template.html", "Path to template html file")
	json := flag.String("-json-path", "./sitemap.json", "Path to sitemap json file")
	cpPath := flag.String("-cp", "./colorpalette.json", "Path to colorpalette json file")
	handler := personalwebsite.HandleSitemap(*html, *json, *cpPath)
	http.ListenAndServe(":20", handler)
}
