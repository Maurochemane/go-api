package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

var (
	host     = getEnv("DB_HOST", "localhost")
	port     = getEnvInt("DB_PORT", 5432)
	user     = getEnv("DB_USER", "postgres")
	password = getEnv("DB_PASSWORD", "0220")
	dbname   = getEnv("DB_NAME", "postgres")
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	// Retry 10 vezes com intervalo de 2 segundos
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Printf("Tentativa %d: Erro ao abrir conexão: %v\n", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			fmt.Println("Connected to " + dbname)

			// Executar migrações automaticamente
			if err := RunMigrations(db); err != nil {
				fmt.Printf("Erro ao rodar migrações: %v\n", err)
				return nil, err
			}

			return db, nil
		}

		fmt.Printf("Tentativa %d: Banco não respondeu: %v\n", i+1, err)
		db.Close()
		time.Sleep(2 * time.Second)
	}

	panic(err)
}
