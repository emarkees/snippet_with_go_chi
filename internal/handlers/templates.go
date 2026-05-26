package handlers

import (
	// "html/template"
	// "path/filepath"

	"github.com/emarkees/chi/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
