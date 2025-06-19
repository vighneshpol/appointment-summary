package database

import (
	"AppointmentSummmary_Assignment/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	cfg := config.GetDBConfig()

	dsn := cfg.GetDSN() // get DSN string
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf(" Failed to connect to database: %v", err)
	}

	DB = db // ✅ Assign to global DB variable
	fmt.Println("✅ Connected to PostgreSQL database successfully.")
}
