package main

import (
	"database/sql"
	"flag"
	"os"

	_ "github.com/lib/pq"
	"github.com/pterm/pterm"
)

var db1 sql.DB
var create, drop, insert *bool

func main() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromString("PSQL")).Render()

	if isConected, _ := initConnection(); !isConected {
		pterm.Error.Println("Erro ao conectar com o banco.")
		os.Exit(1)
	}

	setFlags()

	createTable()
	dropTable()
	insertData()
}

func setFlags() {
	create = flag.Bool("create", false, "create table?")
	drop = flag.Bool("drop", false, "drop table?")
	insert = flag.Bool("insert", false, "insert?")

	flag.Parse()
}

func initConnection() (bool, error) {
	db, err := createConnection()
	defer db.Close()

	if err != nil {
		return false, err
	}

	connectivity := db.Ping()
	if connectivity != nil {
		return false, err
	}

	return true, nil
}

func createConnection() (*sql.DB, error) {
	return sql.Open("postgres", `user=postgres password=changeme 
		host=127.0.0.1 port=5432 dbname=postgres sslmode=disable`)
}

func dropTable() {
	if !*drop {
		pterm.Info.Println("Drop não executado")
		return
	}

	db, err := createConnection()
	defer db.Close()

	DBDrop := `DROP TABLE public.test`
	_, err = db.Exec(DBDrop)
	if err != nil {
		pterm.Error.Println("Erro ao deletar a tabela")
	} else {
		pterm.Success.Println("Tabela deletada com sucesso!")
	}
}

func createTable() {
	if !*create {
		pterm.Info.Println("Create não executado")
		return
	}

	db, err := createConnection()
	defer db.Close()

	DBCreate := ` 
	CREATE TABLE public.test 
	( 
		id integer, 
		name character varying COLLATE pg_catalog."default" 
	) 
	WITH ( 
		OIDS = FALSE 
	)
	TABLESPACE pg_default;
	ALTER TABLE public.test
		OWNER to postgres;
	`
	_, err = db.Exec(DBCreate)
	if err != nil {
		pterm.Error.Println("Erro ao criar a tabela")
	} else {
		pterm.Success.Println("Tabela criada com sucesso!")
	}
}

func insertData() {
	if !*insert {
		pterm.Info.Println("Insert não executado")
		return
	}

	db, err := createConnection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO test VALUES ($1, $2)")
	if err != nil {
		pterm.Error.Println("Não foi possível preparar a inserção!")
	}
	_, err = insert.Exec(2, "second")
	if err != nil {
		pterm.Error.Println("Não foi possível executar a inserção!")
	} else {
		pterm.Success.Println("Registro inserido com sucesso!")
	}
}
