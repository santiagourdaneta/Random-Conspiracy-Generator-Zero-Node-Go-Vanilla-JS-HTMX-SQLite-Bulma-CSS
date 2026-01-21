package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SEOController struct {
	DB *sql.DB
}

// Sitemap devuelve un XML dinámico con todas las teorías
func (sc *SEOController) GetSitemap(c echo.Context) error {
	var slugs []string
	rows, err := sc.DB.Query("SELECT slug FROM theories")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var slug string
		rows.Scan(&slug)
		slugs = append(slugs, slug)
	}

	// Construcción del XML (KISS: strings.Builder es ultra rápido)
	host := c.Request().Host
	now := time.Now().Format("2006-01-02")

	xml := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>http://` + host + `/</loc>
		<lastmod>` + now + `</lastmod>
		<changefreq>daily</changefreq>
		<priority>1.0</priority>
	</url>`

	for _, slug := range slugs {
		xml += fmt.Sprintf(`
	<url>
		<loc>http://%s/t/%s</loc>
		<lastmod>%s</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.7</priority>
	</url>`, host, slug, now)
	}

	xml += "\n</urlset>"

	// Es vital enviar el Content-Type correcto para que Google lo reconozca
	return c.XMLBlob(http.StatusOK, []byte(xml))
}
