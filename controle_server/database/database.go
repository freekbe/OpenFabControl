package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq" // postgress specific package
	"fmt"
	"log"
	"os"
	"time"
)

var Self *sql.DB

func Initdb() {
	pgUser, pgPass, pgHost, pgPort, pgDB := getenv()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", pgUser, pgPass, pgHost, pgPort, pgDB)
	// disable SSL for local development
	dsn += "?sslmode=disable"

	// connect to database
	var err error
	Self, err = connectWithRetries(dsn, 15, 2*time.Second)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// ensure table exists
	if err := ensureTable(); err != nil {
		log.Fatalf("failed to ensure table: %v", err)
	}
}

func getenv() (string, string, string, string, string) {
	pgUser	:= os.Getenv("POSTGRES_USER")
	pgPass	:= os.Getenv("POSTGRES_PASSWORD")
	pgHost	:= os.Getenv("POSTGRES_HOST")
	pgPort	:= os.Getenv("POSTGRES_PORT")
	pgDB	:= os.Getenv("POSTGRES_DB")

	return pgUser, pgPass, pgHost, pgPort, pgDB
}

func connectWithRetries(dsn string, maxRetries int, delay time.Duration) (*sql.DB, error) {
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("db open error: %v", err)
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if pingErr := db.PingContext(ctx); pingErr == nil {
				return db, nil
			} else {
				log.Printf("db ping error: %v", pingErr)
				db.Close()
			}
		}
		log.Printf("retrying DB connection in %s... (%d/%d)", delay, i+1, maxRetries)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("could not connect to DB after %d attempts: last error: %v", maxRetries, err)
}

func ensureTable() error {
	create := `CREATE TABLE IF NOT EXISTS machine_controller (
		id SERIAL PRIMARY KEY,
		uuid TEXT UNIQUE NOT NULL,
		approved BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
	);`
	_, err := Self.Exec(create)
	return err
}
