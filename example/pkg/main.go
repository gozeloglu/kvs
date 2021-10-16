package main

import (
	"fmt"
	"github.com/gozeloglu/kvs"
	"log"
)

func main() {
	db, err := kvs.Create(":1234", "users")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("DB Created.")
	db.Set("john", "23")
	db.Set("jack", "43")

	john := db.Get("john")
	fmt.Println("John:", john)

	jack := db.Get("jack")
	fmt.Println("Jack", jack)

	err = db.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
