package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		// Keep this on a private/admin port or behind auth.
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Start the real service here.
	select {}
}
