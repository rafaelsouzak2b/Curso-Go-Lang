package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria") //Passando parâmetros e executando
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected() //Retorna quantas linhas foram afetadas
	fmt.Println(linhas)
}
