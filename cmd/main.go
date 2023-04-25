package main

import (
	_ "github.com/joho/godotenv/autoload"
	"hte-danger-zone-ms/internal/router"
	"hte-danger-zone-ms/metrics"
	"log"
)

func main() {
	metrics.StartServer()

	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
