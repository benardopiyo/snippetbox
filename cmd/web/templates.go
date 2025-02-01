package main

import "snips/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []models.Snippet
}

