# 1. Validar que la Home carga con el tema Dark Sci-Fi
GET http://localhost:1323/
HTTP 200
[Asserts]
header "Content-Security-Policy" contains "default-src 'self'"
body contains "ARCHIVOS CLASIFICADOS"
body contains "EXTRAER VERDAD"

# 2. Validar que el botón de pánico genera una teoría vía HTMX (POST)
POST http://localhost:1323/generate
HTTP 200
[Asserts]
body contains "teoria-"
body contains "box notification"
body contains "Copiar Enlace"

# 3. Validar el SEO: El Sitemap debe ser XML válido y no estar vacío
GET http://localhost:1323/sitemap.xml
HTTP 200
[Asserts]
header "Content-Type" contains "xml"
body contains "<loc>"
body contains "teoria-"