package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pterm/pterm"
)

var db1 sql.DB

func main() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromString("Postgresql teste")).Render()

	if isConected, _ := initConnection(); !isConected {
		fmt.Println()
		os.Exit(1)
	}

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
		panic(err)
	} else {
		fmt.Println("The table was successfully created!")
	}
}

func initConnection() (bool, error) {
	db, err := sql.Open("postgres", `user=postgres password=changeme 
		host=127.0.0.1 port=5432 dbname=postgres sslmode=disable`)
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
