package models

import (
	"fmt"
	"net/url"
	"strings"
)

type Image struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

// Validate checks if the image data is valid
func (img *Image) Validate() error {
	if strings.TrimSpace(img.Title) == "" {
		return fmt.Errorf("title is required")
	}

	if strings.TrimSpace(img.URL) == "" {
		return fmt.Errorf("URL is required")
	}

	// Validate URL format
	if _, err := url.ParseRequestURI(img.URL); err != nil {
		return fmt.Errorf("invalid URL format")
	}

	// Validate alt text quality (for extension tasks)
	if err := img.validateAltText(); err != nil {
		return err
	}

	return nil
}

func (img *Image) validateAltText() error {
	altText := strings.TrimSpace(strings.ToLower(img.AltText))

	// Check for useless alt text
	uselessTexts := []string{"image", "photo", "picture", "img"}
	for _, useless := range uselessTexts {
		if altText == useless {
			return fmt.Errorf("alt text too generic: '%s'", img.AltText)
		}
	}

	if len(altText) < 5 {
		return fmt.Errorf("alt text too short, should be descriptive")
	}

	return nil
}
