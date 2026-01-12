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
	// create machine controler table
	create := `CREATE TABLE IF NOT EXISTS machine_controller (
		id SERIAL PRIMARY KEY,
		uuid TEXT UNIQUE NOT NULL,
		type TEXT NOT NULL,
		zone TEXT NOT NULL,
		name TEXT NOT NULL,
		manual TEXT NOT NULL,
		price_booking_in_eur FLOAT NOT NULL,
		price_usage_in_eur FLOAT NOT NULL,
		approved BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
	);` // FLOAT is synonym to double pressision (64 bit float)
	_, err := Self.Exec(create)
	if err != nil { return err }

	// create the roles table
	create = `CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		privilege_lvl INTEGER,
		name TEXT
	);`
	_, err = Self.Exec(create)
	if err != nil { return err }

	// create users table
	create = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		access_key TEXT NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255),
		first_name VARCHAR(64),
		last_name VARCHAR(64),
		tva VARCHAR(16),
		facturation_address VARCHAR(255),
		account VARCHAR(34),
		verification_code VARCHAR(32) NOT NULL,
		status VARCHAR(16) DEFAULT 'pending',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
	);`
	_, err = Self.Exec(create)
	if err != nil { return err }

	// create user_roles junction table
	create = `CREATE TABLE IF NOT EXISTS user_roles (
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, role_id)
	);`
	_, err = Self.Exec(create)
	if err != nil { return err }

	// create user_machines junction table
	create = `CREATE TABLE IF NOT EXISTS user_machines (
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		machine_id INTEGER REFERENCES machine_controller(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, machine_id)
	);`
	_, err = Self.Exec(create)
	if err != nil { return err }

	return nil
}
