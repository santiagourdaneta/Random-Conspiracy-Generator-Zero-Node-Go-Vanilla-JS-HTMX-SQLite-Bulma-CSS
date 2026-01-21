BINARY=conspiracy_ai.exe

.PHONY: all dev build test stress e2e clean

# Modo desarrollo con Hot Reload
dev:
	air

all: build test e2e

# Compilación para Windows (Laptops viejas: -s -w reduce el tamaño un 30%)
build:
	go mod tidy
	go build -ldflags="-s -w" -o $(BINARY) main.go

# Tests de lógica MVC
test:
	go test ./models/... ./controllers/... -v

# Pruebas de integración E2E con Hurl
e2e:
	@echo "👁️ Verificando integridad del sistema..."
	hurl --test tests/e2e.hurl

# Prueba de carga (Requiere Hey instalado)
stress:
	@echo "🔥 Estresando el servidor (1000 requests)..."
	hey -n 1000 -c 100 -m POST http://localhost:1323/generate

clean:
	rm -f $(BINARY)
	rm -f conspiracies.db
