package controllers

import (
	"conspiracy-app/database"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateEndpoint(t *testing.T) {
	// 1. Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/generate", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// DB en memoria para el test
	db := database.InitDB(":memory:")
	database.PopulateIfEmpty(db)
	tc := &TheoryController{DB: db}

	// 2. Ejecución
	if assert.NoError(t, tc.Generate(c)) {
		// 3. Aserciones
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "is-size-5") // Verifica que hay un párrafo de teoría
		assert.Contains(t, rec.Body.String(), "🔎")         // Verifica nuestro nuevo icono
	}
}
