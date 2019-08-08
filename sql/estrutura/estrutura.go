package main

//O driver do mysql deve estar instalado(go get -u github.com/go-sql-driver/mysql)

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //uso implícito
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	//db, err := sql.Open("mysql", "usuario:senha@banco de dados")
	db, err := sql.Open("mysql", "root@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
