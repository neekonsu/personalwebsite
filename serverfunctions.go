package personalwebsite

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

// HandleSitemap generates an http.HandlerFunc given a Sitemap and path to an html template
func HandleSitemap(html string, json string, cpPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sitemap := *MakeSitemap(json, cpPath)
		if page, exists := sitemap[r.URL.Path[1:]]; exists {
			colorpalette := MakeColorPalette(cpPath)
			page.PopulateColors(colorpalette)
			MakeTemplate(html).Execute(w, page)
		} else if r.URL.Path == "" || r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusFound)
		} else if r.URL.Path == "/static/stylesheet.css" {
			dir := http.Dir("./res/")
			file, err := dir.Open("stylesheet.css")
			Check(err)
			defer file.Close()
			fi, err := file.Stat()
			if err != nil {
				log.Printf("HTTP error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text/css")
			http.ServeContent(w, r, r.URL.Path, fi.ModTime(), file)
		} else {
			http.Redirect(w, r, "/pagenotfound", http.StatusFound)
		}
	}
}

// MakeTemplate generates *template.Template from the path to an html template
func MakeTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(path))
}

// MakeSitemap takes path to the sitemap json file and returns a populated Sitemap struct
func MakeSitemap(path string, cpPath string) *Sitemap {
	output := make(Sitemap)
	reader, err := os.Open(path)
	Check(err)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&output)
	return &output
}

// MakeColorPalette same as MkeSitemap but for the type ColorPalette
func MakeColorPalette(path string) ColorPalette {
	var output ColorPalette
	reader, err := os.Open(path)
	Check(err)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&output)
	return output
}

// Check checks for errors :)
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
