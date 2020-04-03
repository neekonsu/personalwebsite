package personalwebsite

import (
	"math/rand"
	"time"
)

// ColorPalette represents all the colors that the site will take
type ColorPalette struct {
	Colors []string `json:"colors,omitempty"`
}

// Sitemap represents the relative path to each page with the key and the value is the page's struct
type Sitemap map[string]Page

// Page stores the title of the page along with all the cards to be rendered
type Page struct {
	Title string `json:"title,omitempty"`
	Cards []Card `json:"cards,omitempty"`
}

// Card represents a single div on any individual page and stores all necessary information for that card
type Card struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"paragraphs,omitempty"` // each paragraph will be represented by one string in order to be rendered as a single paragraph; no need for inherent formatting
	URL        string   `json:"url,omitempty"`
	Color      string   `json:"color,omitempty"`
}

// PickColor returns a single random color from the colorpalette
func (c *ColorPalette) PickColor() string {
	return c.Colors[rand.Intn(len(c.Colors))]
}

// NewColor picks a random color other than the current color
func (c *Card) NewColor(palette ColorPalette, not ...[]string) {
	tempColor := palette.PickColor()
	for { // this method will limit the number of sections to the length of the color palette;
		// if you have equal to or greater than the number of colors in sections, then this loop will stack overflow
		collision := false
		for _, colorSlice := range not {
			for _, color := range colorSlice {
				if color == tempColor {
					tempColor = palette.PickColor()
					collision = true
				}
			}
		}
		if !collision {
			c.Color = tempColor
			break
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}
