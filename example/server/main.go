package main

import (
	"github.com/gozeloglu/kvs"
	"log"
)

func main() {
	db, err := kvs.Create(":1234", "users")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("DB Created.")
	db.Open()
}
