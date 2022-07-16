package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paganotoni/tato/migrations"
	"github.com/paganotoni/tato/web"
)

func main() {
	pwd, _ := os.Getwd()
	dsn := filepath.Join(pwd, "tato.db")

	db, err := migrations.Run(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("> Tato Server running on: http://localhost:3000")
	err = http.ListenAndServe(":3000", web.Server(db))
	if err != nil {
		log.Fatal(err)
	}
}
