# 👁️ Random-Conspiracy-Generator

[![Go Quality Control](https://github.com/santiagourdaneta/Random-Conspiracy-Generator-Zero-Node-Go-Vanilla-JS-HTMX-SQLite-Bulma-CSS/actions/workflows/ci.yml/badge.svg)](https://github.com/santiagourdaneta/Random-Conspiracy-Generator-Zero-Node-Go-Vanilla-JS-HTMX-SQLite-Bulma-CSS/actions)

Esta es una aplicación web **Full-Stack** inspirada en la ética Punk y la complejidad del Mathcore. Diseñada para hardware antiguo, priorizando la velocidad de ejecución y la simplicidad arquitectónica.

## 🚀 Características (Zero-Node Philosophy)
- **Backend:** Go 1.21+ (Compilado, concurrente, < 20MB RAM).
- **Frontend:** HTML-First con **HTMX** (Sin JavaScript pesado).
- **Estética:** Dark Mode Sci-Fi / Illuminati con **Bulma CSS**.
- **Base de Datos:** SQLite3 (Embebida, sin necesidad de servidores externos).
- **PWA:** Funciona offline gracias a Service Workers nativos.
- **Seguridad:** CSP (Content Security Policy) y protección CSRF.

## 🛠️ Stack Tecnológico
- **Lenguaje:** Go (Golang)
- **Web Framework:** Echo (V4)
- **UI:** Bulma CSS + Animate.css
- **Interactividad:** HTMX 
- **Persistencia:** SQLite3

## 📦 Instalación y Uso

### Requisitos
- Go 1.21 o superior.
- GCC (para el driver de SQLite3).

### Pasos
1. Clonar el repositorio:
   ```bash
   git clone https://github.com/santiagourdaneta/Random-Conspiracy-Generator-Zero-Node-Go-Vanilla-JS-HTMX-SQLite-Bulma-CSS](https://github.com/santiagourdaneta/Random-Conspiracy-Generator-Zero-Node-Go-Vanilla-JS-HTMX-SQLite-Bulma-CSS)
   cd Random-Conspiracy-Generator-Zero-Node-Go-Vanilla-JS-HTMX-SQLite-Bulma-CSS
   
2. Instalar dependencias y limpiar módulos:

go mod tidy

3. Ejecutar en modo desarrollo:

make dev

🧪 Testing y Calidad

make test      # Unit & Integration Tests
make e2e       # End-to-End con Hurl
make stress    # Stress testing con Hey

📐 Arquitectura MVC

El proyecto sigue una separación estricta de responsabilidades:

models/: Estructuras de datos.

views/: Plantillas HTML puras.

controllers/: Lógica de negocio y manejo de peticiones.

database/: Capa de persistencia y seeding.

Hecho con furia y precisión. Inspirado en el caos técnico de The Fall of Troy.
   
