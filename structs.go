package personalwebsite

import (
	"math/rand"
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
func (c *ColorPalette) PickColor() (string, int) {
	// rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	index := rand.Intn(len(c.Colors))
	output := c.Colors[index]
	return output, index
}

// PopulateColors picks a random color other than the current color
func (p *Page) PopulateColors(palette ColorPalette) {
	cardColors := make([]*string, len(p.Cards))
	for i := range p.Cards {
		cardColors = append(cardColors, &p.Cards[i].Color)
	}
	for i := range cardColors {
		color, index := palette.PickColor()
		if pointer := cardColors[i]; pointer != nil {
			*pointer = color
			if index != len(palette.Colors)-1 {
				palette.Colors = append(palette.Colors[:index], palette.Colors[index+1:]...)
			} else {
				palette.Colors = palette.Colors[:index]
			}
		}
	}
}
