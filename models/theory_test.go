package models

import (
	"strings"
	"testing"
)

func TestTheoryValidation(t *testing.T) {
	// Caso de prueba: Una teoría válida
	theory := Theory{
		Content: "Los delfines controlan el 5G desde la Antártida.",
		Slug:    "teoria-delfines-5g",
	}

	if theory.Content == "" {
		t.Error("ERROR: El contenido de la teoría no puede estar vacío")
	}

	if !strings.HasPrefix(theory.Slug, "teoria-") {
		t.Errorf("ERROR: El slug '%s' debe empezar con el prefijo 'teoria-'", theory.Slug)
	}
}
