@echo off
echo [ACCESO NIVEL 7] Preparando sistema...
go mod tidy
echo [INFO] Compilando binario de alto rendimiento...
go build -ldflags="-s -w" -o conspiracy_ai.exe main.go
if %errorlevel% neq 0 (
    echo [ERROR] Fallo en la compilacion. Revisa los logs.
    pause
    exit /b
)
echo [OK] Binario generado: conspiracy_ai.exe
echo [INFO] Ejecutando tests de integridad...
make test
echo [OK] Iniciando servidor en http://localhost:1323
start http://localhost:1323
conspiracy_ai.exe
pause