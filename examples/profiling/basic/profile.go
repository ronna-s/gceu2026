package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:6060", nil))
}
