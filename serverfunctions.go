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
		sitemap := MakeSitemap(json, cpPath)
		if page, exists := sitemap[r.URL.Path]; exists {
			MakeTemplate(html).Execute(w, page)
		} else if r.URL.Path == "" || r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusFound)
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
func MakeSitemap(path string, cpPath string) Sitemap {
	output := make(Sitemap)
	colorpalette := MakeColorPalette(cpPath)
	reader, err := os.Open(path)
	Check(err)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&output)
	for _, page := range output {
		var notColorList []string
		for _, card := range page.Cards {
			card.NewColor(colorpalette, notColorList)
		}
	}
	return output
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
