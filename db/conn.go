package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importação anônima do driver PostgreSQL
)

const (
	host     = "localhost" // Corrigido para "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234" // Adicione a senha do usuário
	dbname   = "postgres"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar a conexão com o banco de dados: %w", err)
	}
	fmt.Println("Conexão com o banco de dados", dbname, "foi estabelecida com sucesso.")
	return db, nil
}
