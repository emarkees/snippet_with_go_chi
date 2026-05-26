package app

import (
	"log"
	"html/template"
	"github.com/emarkees/chi/internal/models"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Snippets *models.SnippetModel
	TemplateCache map[string]*template.Template
}
