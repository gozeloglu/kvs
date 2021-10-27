package main

import (
	"fmt"
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
	db.Set("john", "23")
	db.Set("jack", "43")

	john := db.Get("john")
	fmt.Println("John:", john)

	jack := db.Get("jack")
	fmt.Println("Jack:", jack)

	db.Del("jack")

	jack = db.Get("jack")
	fmt.Println("Jack:", jack)

	newAge, err := db.Incr("john")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("John's new age is %s\n", newAge)

	newAge, err = db.IncrBy("john", 3)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("John's new age is %s\n", newAge)

	err = db.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
