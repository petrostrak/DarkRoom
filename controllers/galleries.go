package controllers

import (
	"DarkRoom/models"
	"DarkRoom/views"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

// The Users structure
type Galleries struct {
	New *views.View
	gs  models.GalleryService
}
