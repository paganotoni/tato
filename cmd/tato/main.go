package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paganotoni/tato/migrations"
	"github.com/paganotoni/tato/storage"
	"github.com/paganotoni/tato/web"
)

func main() {
	pwd, _ := os.Getwd()
	db, err := sql.Open("sqlite3", filepath.Join(pwd, "tato.db"))
	if err != nil {
		log.Fatal(err)
	}

	storage.DB = db
	defer storage.DB.Close()

	err = migrations.Run(storage.DB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("> Tato Server running on: http://localhost:3000")
	err = http.ListenAndServe(":3000", web.Server())
	if err != nil {
		log.Fatal(err)
	}
}
