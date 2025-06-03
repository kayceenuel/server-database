package main

import (
	"net/http"
)

// image struct represents data about an image metadata
type Image struct {
	Title   string `json: "title"`
	AltText string `json: "alt_text"`
	URL     string `json: "url"`
}

// initialising image data 
images := []Image{
    {
      title: "Sunset",
      alt_text: "Clouds at sunset",
      url: "https://images.unsplash.com/photo-1506815444479-bfdb1e96c566?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
    },
    {
      title: "Mountain",
      alt_text: "A mountain at sunset",
      url: "https://images.unsplash.com/photo-1540979388789-6cee28a1cdc9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
    },
}

func main() {
	
}