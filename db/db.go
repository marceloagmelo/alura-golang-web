package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//ConectaComBancoDeDados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin123 host=192.168.15.200 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
