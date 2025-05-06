package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jcompagni10/skip-router-data/x/reporter"
	"github.com/jcompagni10/skip-router-data/x/skip"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func initDB(dbPath string) (*sql.DB, error) {
	log.Printf("Initializing database at %s", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Create table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS swap_routes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		winner TEXT,
		winning_price REAL,
		neutron_price REAL,
		token_in TEXT,
		token_out TEXT,
		amount_in INTEGER,
		time DATETIME,
		source_chain TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	log.Info("Database initialized successfully")
	return db, nil
}

func SetupDB() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/database.db"
	}

	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	db, err := initDB(dbPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	return db, nil
}

func SetupData() {
	reporter.ParsePairsFromEnv()
	reporter.ParseChainIdsFromEnv()

	chainAssets, err := skip.GetChainAssets(reporter.ChainIds)
	if err != nil {
		log.Fatalf("Error getting chain assets: %v", err)
	}
	reporter.ChainData = chainAssets

	reporter.ParseExclusionsFromEnv()

	symbolsToSeed := []string{}
	for _, pair := range reporter.Pairs {
		symbolsToSeed = append(symbolsToSeed, pair[1])
	}

	reporter.SeedPriceCache(symbolsToSeed)
}

func main() {
	var sleepTime = 30 * time.Second
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting application...")

	db, err := SetupDB()
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	defer db.Close()

	SetupData()

	for {
		log.Info("Running ReportSwapRoutes...")
		reporter.ReportSwapRoutes(db)
		log.Infof("Sleeping for %v...", sleepTime)
		time.Sleep(sleepTime)
	}
}
