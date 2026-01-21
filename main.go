package main

import (
	"log/slog"
	"net/http"
	"os"

	"conspiracy-app/controllers"
	"conspiracy-app/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	// 1. OBSERVABILIDAD: Logs Estructurados JSON (Zero Node / Zero Heavy Tools)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. INFRAESTRUCTURA: Inicializar DB SQLite
	db := database.InitDB("./conspiracies.db")
	database.PopulateIfEmpty(db)
	defer db.Close()

	// 3. CONTROLADORES
	theoryController := &controllers.TheoryController{DB: db}
	seoController := &controllers.SEOController{DB: db}

	e := echo.New()

	// 4. MIDDLEWARES DE SEGURIDAD & RENDIMIENTO
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Logger de peticiones en JSON
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true, LogURI: true, LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info("request", "method", v.Method, "uri", v.URI, "status", v.Status, "ip", c.RealIP())
			return nil
		},
	}))

	// BLINDAJE CSP: Bloquea cualquier script inyectado externamente
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' https://unpkg.com; " +
			"style-src 'self' 'unsafe-inline' https://cdn.jsdelivr.net; " +
			"font-src 'self' data:; " +
			"img-src 'self' data:; " +
			"connect-src 'self';",
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "DENY",
		HSTSMaxAge:         31536000,
	}))

	// Rate Limiting: 10 peticiones por segundo por IP
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// 5. RUTAS (Zero Node - SSR & HTMX)

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		// Esto busca el archivo relativo a donde ejecutas el programa
		return c.File("views/index.html")
	})

	e.GET("/t/:slug", theoryController.ViewTheory)

	// API interna para HTMX (Sin recarga de página)
	e.POST("/generate", theoryController.Generate)

	// SEO Pro: Sitemap dinámico con Caching
	e.GET("/sitemap.xml", seoController.GetSitemap)

	// HONEYPOT: Atrapa bots intentando hackear
	e.GET("/admin-panel", func(c echo.Context) error {
		slog.Warn("BOT_DETECTED", "ip", c.RealIP(), "path", c.Path())
		return c.String(http.StatusTeapot, "Access Denied: You have been reported to the Order.")
	})

	// 6. LANZAMIENTO
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	slog.Info("Server started", "port", port, "env", "production")
	e.Logger.Fatal(e.Start(":" + port))
}
