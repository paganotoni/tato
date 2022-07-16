package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paganotoni/tato/web"
)

func main() {
	fmt.Println("> Tato Server running on: http://localhost:3000")
	err := http.ListenAndServe("0.0.0.0:3000", web.Server())
	if err != nil {
		log.Fatal(err)
	}
}
