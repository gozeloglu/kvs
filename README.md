# kvs

## Installation

You can add package to your project with the following command.

```shell
go get github.com/gozeloglu/kvs
```

## Example

```go
package main

import "github.com/gozeloglu/kvs"

func main() {
    db, err := kvs.Open("users")
    if err != nil {
        log.Fatalf(err.Error())
    }
    defer db.Close()
}
```
