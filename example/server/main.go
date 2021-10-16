package main

import (
	"github.com/gozeloglu/kvs"
	"log"
	"time"
)

func main() {
	db, err := kvs.Create(":1234", "users", 1*time.Minute)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("DB Created.")
	db.Open()
}
