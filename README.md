# kvs

## Installation

You can add package to your project with the following command.

```shell
go get github.com/gozeloglu/kvs
```

## Example

```go
package main

import (
	"fmt"
	"github.com/gozeloglu/kvs"
	"log"
)

func main() {
    db, err := kvs.Open("users")
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

## LICENSE

[MIT](https://github.com/gozeloglu/kvs/blob/main/LICENSE)
