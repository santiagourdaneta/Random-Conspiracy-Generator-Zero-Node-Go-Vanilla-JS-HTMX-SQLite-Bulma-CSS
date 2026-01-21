package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2" // API moderna de Go 1.22+

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal("❌ Error al abrir la DB:", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS theories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("❌ Error al crear tablas:", err)
	}

	return db
}

func PopulateIfEmpty(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM theories").Scan(&count)
	if count > 0 {
		return
	}

	log.Println("📜 Generando 500 archivos clasificados...")

	sujetos := []string{"El Vaticano", "La élite financiera", "La inteligencia artificial", "La NASA", "El Nuevo Orden Mundial"}
	verbos := []string{"está manipulando", "ha interceptado", "está decodificando", "ha clonado en secreto", "está transmitiendo"}
	objetos := []string{"las señales de radio de la Antártida", "el ADN de la población civil", "las frecuencias del núcleo terrestre", "los archivos perdidos de Tesla"}
	complementos := []string{"para instaurar un control total.", "bajo órdenes de entidades no humanas.", "mediante el uso de satélites ocultos.", "antes del evento del Gran Reinicio."}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// rand.N(len) es la forma moderna y segura de obtener un índice aleatorio
	for i := range 500 {
		content := fmt.Sprintf("%s %s %s %s",
			sujetos[rand.N(len(sujetos))],
			verbos[rand.N(len(verbos))],
			objetos[rand.N(len(objetos))],
			complementos[rand.N(len(complementos))])

		slug := fmt.Sprintf("veritas-%d-%x", i+1, rand.Uint32())
		_, err := tx.Exec("INSERT INTO theories (content, slug) VALUES (?, ?)", content, slug)
		if err != nil {
			log.Printf("⚠️ Error insertando teoría %d: %v", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("❌ Error al confirmar la transacción:", err)
	}

	log.Println("✅ Base de datos poblada con éxito.")
}
