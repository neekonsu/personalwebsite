package main

import (
	"flag"
	"net/http"

	"github.com/neekonsu/personalwebsite"
)

func main() {
	html := flag.String("-html-path", "./res/template.html", "Path to template html file")
	json := flag.String("-json-path", "./res/portfolio.json", "Path to sitemap json file")
	cpPath := flag.String("-cp", "./res/colorpalette.json", "Path to colorpalette json file")
	handler := personalwebsite.HandleSitemap(*html, *json, *cpPath)
	http.ListenAndServe(":8080", handler)
}
