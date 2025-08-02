package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=public sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println("Erro ao abrir conexão:", err)
		} else {
			err = db.Ping()
			if err == nil {
				fmt.Println("Conectado ao banco:", dbname)
				return db, nil
			}
			fmt.Println("Erro no ping:", err)
		}
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("não conseguiu conectar ao banco após várias tentativas: %w", err)
}
