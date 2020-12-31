package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/pterm/pterm"
)

var db1 sql.DB

func main() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromString("PSQL")).Render()

	if isConected, _ := initConnection(); !isConected {
		pterm.Error.Println("Erro ao conectar com o banco.")
		os.Exit(1)
	}

	createTable()

	db, err := sql.Open("postgres", `user=postgres password=changeme 
		host=127.0.0.1 port=5432 dbname=postgres sslmode=disable`)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	connectivity := db.Ping()
	if connectivity != nil {
		panic(err)
	}

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

func createTable() {
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
