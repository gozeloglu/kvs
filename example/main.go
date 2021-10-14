package example

import (
	"github.com/gozeloglu/kvs"
	"log"
)

func main() {
	db, err := kvs.Open("users")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()
}
