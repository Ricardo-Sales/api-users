package main

import (
	"log"
	"net/http"

	"github.com/Ricardo-Sales/api-users/routers"
)

func main() {

	r := routers.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))

}
