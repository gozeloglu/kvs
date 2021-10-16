# kvs

## Installation

You can add package to your project with the following command.

```shell
go get github.com/gozeloglu/kvs
```

## Example

If you want to use in your code as a package, you can call `Get` and `Set` methods directly.

```go
package main

import (
    "fmt"
    "github.com/gozeloglu/kvs"
    "log"
)

func main() {
    db, err := kvs.Open("", "users")
    if err != nil {
        log.Fatalf(err.Error())
    }
    
    db.Set("john", "23")
    db.Set("jack", "43")
    
    john := db.Get("john")
    fmt.Println(john)
    
    jack := db.Get("jack")
    fmt.Println(jack)
    
    err = db.Close() // Call while closing the database.
    if err != nil {
        log.Fatalf(err.Error())
    }
}

```

If you want to use as a server, you can just call two different functions. It creates endpoints for you.

```go
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

```

## LICENSE

[MIT](https://github.com/gozeloglu/kvs/blob/main/LICENSE)
