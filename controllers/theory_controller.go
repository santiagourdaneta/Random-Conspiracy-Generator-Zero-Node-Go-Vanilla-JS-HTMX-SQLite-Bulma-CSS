package controllers

import (
	"conspiracy-app/models"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type TheoryController struct {
	DB *sql.DB
}

func (tc *TheoryController) Generate(c echo.Context) error {
	var t models.Theory
	err := tc.DB.QueryRow("SELECT content, slug FROM theories ORDER BY RANDOM() LIMIT 1").Scan(&t.Content, &t.Slug)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error en la Matrix")
	}

	// Usamos un tag <a> con estilo de botón pequeño
	html := fmt.Sprintf(`
			<div class="box notification is-primary is-light animate__animated animate__fadeIn">
				<p class="is-size-5">🔎 %s</p>

			</div>`, t.Content)

	return c.HTML(http.StatusOK, html)

}

func (tc *TheoryController) ViewTheory(c echo.Context) error {
	slug := c.Param("slug")
	var t models.Theory

	err := tc.DB.QueryRow("SELECT content, slug FROM theories WHERE slug = ?", slug).Scan(&t.Content, &t.Slug)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	dat, err := os.ReadFile("views/index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: Plantilla no encontrada")
	}

	htmlTemplate := string(dat)

	contentHTML := fmt.Sprintf(`
		<div class="box notification is-primary is-light">
			<p class="is-size-5">🔎 %s</p>
			<hr>
			<a href="/" class="button is-small is-dark">Volver al Generador</a>
		</div>`, t.Content)

	finalHTML := strings.Replace(htmlTemplate, "", contentHTML, 1)

	return c.HTML(http.StatusOK, finalHTML)
}
