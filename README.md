# kvs [![Go Reference](https://pkg.go.dev/badge/github.com/gozeloglu/kvs.svg)](https://pkg.go.dev/github.com/gozeloglu/kvs) [![Go Report Card](https://goreportcard.com/badge/github.com/gozeloglu/kvs)](https://goreportcard.com/report/github.com/gozeloglu/kvs) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gozeloglu/kvs) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/gozeloglu/kvs) [![API Doc](https://img.shields.io/badge/-API%20Doc-orange)](https://github.com/gozeloglu/kvs/wiki/API-Documentation)  [![LICENSE](https://img.shields.io/badge/license-MIT-green)](https://github.com/gozeloglu/kvs/blob/main/LICENSE)

kvs is an in-memory key-value storage written in Go. It has 2 different usage. It can be used as a package by importing
it to your code or as a server by creating an HTTP server. kvs stores persistent data in local machine's `/tmp/kvs`
directory. The file extension is `.kvs`. For example, if you create a database named as **user**, it would be stored in
a file name as `users.kvs`. It loads the data from file into memory if the database is already exists. Also, kvs
supports saving data from memory to disk in a given time interval periodically. You can specify the time interval while
creating the database. Both keys and values are stored as string. That's why the methods accept only strings.

## Installation

You can add package to your project with the following command.

```shell
go get github.com/gozeloglu/kvs
```

## Package Usage

Firstly, you need to create a database by calling `kvs.Open()`. It creates a new database if not exists or loads the
data from existing database if it exists. If you want to use kvs as a package, you don't need to specify `addr` as a
first parameter. As a third parameter, you pass time interval to save data to database periodically.

```go
// Creates a "users" database and saves the data from memory to file per 2 minutes.
db, err := kvs.Open("", "users", 2*time.Minute)  
```

Then, simply you can call `Set()` and `Get()` methods. `Set()` takes key and value as parameters and adds the key-value
pair to memory. `Get()` takes key as a parameter and returns the value of the key. Both `Set()` and `Get()` methods
takes string as parameters.

```go
// "john" is stored as key with "23" as value in memory.
db.Set("john", "23")

// Returns "23" to age.
age := db.Get("john")
```

If you want to make sure that all data stores in memory would save to disk, you can call `Close()` method. It writes the
data to disk and closes the database.

```go
// Writes data in memory to disk
db.Close()
```

If you want to see full code, you can take a
look [/example/pkg/main.go](https://github.com/gozeloglu/kvs/blob/main/example/pkg/main.go).

## Package Example

```go
package main

import (
	"fmt"
	"github.com/gozeloglu/kvs"
	"log"
	"time"
)

func main() {
	db, err := kvs.Open("", "users", 2*time.Minute)
	if err != nil {
		log.Fatalf(err.Error())
	}

	db.Set("john", "23")
	db.Set("jack", "43")

	johnAge := db.Get("john")
	fmt.Println(johnAge)

	jackAge := db.Get("jack")
	fmt.Println(jackAge)

	db.Del("jack")

	jack = db.Get("jack")
	fmt.Println("Jack:", jack)

	newAge, err := db.Incr("john")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("John's new age is %s", newAge)

	newAge, err = db.IncrBy("john", 3)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("John's new age is %s\n", newAge)

	exist := db.Exists("john")
	fmt.Println(exist)

	exist = db.Exists("jack")
	fmt.Println(exist)

	ok, err := db.Rename("john", "john john")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("key name changed: %v1\n", ok)
	
	err = db.Close() // Call while closing the database.
	if err != nil {
		log.Fatalf(err.Error())
	}
}

```

## Server Usage

Server usage is so simple and short. You would call extra method, `Open()`, to start server. Default port is **1234**
for kvs server. But, you can override it and specify another port number.

```go
// Server runs on localhost:1234
db, _ := kvs.Create(":1234", "users", 2*time.Minute)

// The server is started
db.Open()
```

If you want to see full code, you can take a
look [/example/server/main.go](https://github.com/gozeloglu/kvs/blob/main/example/server/main.go). You can run this code
directly without any configurations.

You can find the **API Documentation** from  [**here**](https://github.com/gozeloglu/kvs/wiki/API-Documentation) .

## Server Example

If you want to use as a server, you can just call two different functions. It creates endpoints for you.

```go
package main

import (
	"github.com/gozeloglu/kvs"
	"log"
	"time"
)

func main() {
	db, err := kvs.Create(":1234", "users", 2*time.Minute)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("DB Created.")
	db.Open()
}
```

## How to store data in file?

Key-value pairs store in files with `.kvs` extension. Data format is simple. There is **=** between key and value.

```
foo=bar
john=12
fizz=buzz
```

This is a sample data file.

## NOTE

kvs is still under development stage, and it is created for experimental purposes.

## LICENSE

[MIT](https://github.com/gozeloglu/kvs/blob/main/LICENSE)
