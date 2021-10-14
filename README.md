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
    kvs, err := db.Open("users")
    if err != nil {
        log.Fatalf(err.Error())
    }
    defer kvs.DbFile.Close()
}
```
