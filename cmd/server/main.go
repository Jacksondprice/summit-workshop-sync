package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/summit-workshop-sync/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/schedules", handlers.ListSchedules)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
