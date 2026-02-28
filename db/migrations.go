package db

import (
	"database/sql"
	"fmt"
)

// RunMigrations executa todas as migrações do banco de dados
func RunMigrations(db *sql.DB) error {
	fmt.Println("🔄 Rodando migrações...")

	// Criação da tabela products
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS product (
		id_product SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	createStudebtsTable := `
	CREATE TABLE IF NOT EXISTS student (
		id_student SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`


	// Criar índice para melhor performance
	createIndexOnName := `
	CREATE INDEX IF NOT EXISTS idx_product_name ON product(name);
	`

	// Executar migrations em ordem
	migrations := []string{
		createProductsTable,
		createIndexOnName,
		createStudebtsTable,
	}

	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return fmt.Errorf("erro ao executar migration: %v", err)
		}
	}

	fmt.Println("✅ Migrações executadas com sucesso!")
	return nil
}
