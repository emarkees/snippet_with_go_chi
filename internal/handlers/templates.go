package handlers

import (
	"github.com/emarkees/chi/internal/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}