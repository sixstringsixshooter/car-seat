package main

import (
	"log"
	"net/http"

	"github.com/sixstringsixshooter/car-seat/config"
)

func main() {
	vers := config.Versions{
		"v1": true,
	}
	c, err := newController(&config.Config{Versions: vers})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("starting API...")
	if err := http.ListenAndServe(":8080", c.router); err != nil {
		log.Fatal(err.Error())
	}
}
