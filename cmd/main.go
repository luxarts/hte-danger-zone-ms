package main

import (
	_ "github.com/joho/godotenv/autoload"
	"hte-danger-zone-ms/internal/router"
	"log"
)

func main() {
	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
