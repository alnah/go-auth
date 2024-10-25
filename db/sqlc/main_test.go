package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	cg "github.com/alnah/go-auth/config"
	db "github.com/alnah/go-auth/db/dsn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries
var testDB *sql.DB
var tableNames = []string{"user"}

func TestMain(m *testing.M) {
	var err error

	config, err := cg.LoadConfig("../../.env/", "test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Println(config.PostgresPassword)

	dsn, err := db.GenerateDSN(db.GenerateDSNParams{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		DBName:   config.PostgresName,
	})
	if err != nil {
		log.Fatal("cannot generate DSN:", err)
	}

	testDB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	defer func() {
		if err := testDB.Close(); err != nil {
			log.Fatal("cannot close the database connection: ", err)
		}
	}()

	defer func() {
		if err := cleanupDatabase(testDB); err != nil {
			log.Fatalf("cleanup failed: %v", err)
		}
	}()

	testQueries = New(testDB)
	if testQueries == nil {
		log.Fatal("failed to create testQueries: testDB may be nil")
	}

	test := m.Run()
	os.Exit(test)
}

func cleanupDatabase(connection *sql.DB) error {
	// Retrieve table names from all schemas
	q := "SELECT tablename, schemaname FROM pg_tables"
	rows, err := connection.Query(q)
	if err != nil {
		return fmt.Errorf("cannot retrieve table names: %w", err)
	}

	// Ensure rows are closed after use
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("cannot close rows: %v", err)
		}
	}()

	// Get table names in the database
	var tableNames []string
	for rows.Next() {
		var tableName, schemaName string
		if err := rows.Scan(&tableName, &schemaName); err != nil {
			return fmt.Errorf("cannot scan table name: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}

	// Truncate each table with error handling
	for _, table := range tableNames {
		_, err := connection.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			return fmt.Errorf("cannot clean the database for table %s: %w",
				table, err)
		}
	}
	return nil
}
