package banco

import (
	"database/sql"
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "go_studies"
)

func Conectar() (*sql.DB, error) {
	stringConexao := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, erro := sql.Open("postgres", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
