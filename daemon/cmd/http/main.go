package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"victord/daemon/internal/config"
	api "victord/daemon/transport/http"
)

func main() {
	flag.Parse()

	router := api.SetupRouter()

	address := fmt.Sprintf("%s:%s", *config.Host, strconv.Itoa(*config.Port))

	log.Printf("Victor daemon running on %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}
