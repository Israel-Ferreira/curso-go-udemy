package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	usuarioDb := "golang2"
	senha := "golang2"
	db := "devbook"

	conexaoStrDb := fmt.Sprintf("%s:%s@tcp(%s:33406)/%s?charset=utf8&parseTime=True", usuarioDb, senha, "localhost", db)

	fmt.Println(conexaoStrDb)

	dbConn, err := sql.Open("mysql", conexaoStrDb)

	if err != nil {
		log.Fatalln(err)
	}

	defer dbConn.Close()

	if err = dbConn.Ping(); err != nil {
		log.Fatalln(err)
	}


	fmt.Println("Conexão com o DB Aberta")


	nLin, err := dbConn.Query("SELECT * FROM usuarios")

	if err != nil {
		log.Fatal(err)
	}

	defer nLin.Close()


	fmt.Println(nLin)
}
