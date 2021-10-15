package kvs

import (
	"fmt"
	"log"
	"net/http"
)

// Open creates an HTTP and database connection. HTTP connection listens HTTP
// requests from client and database connection listens database commands from
// server. addr is the HTTP address and dbName is the database name. If HTTP
// address is empty, localhost and default port is used. In contrast, dbName
// needs to be specified. If it is not specified, it returns error and the
// connection is not established.
func Open(addr string, dbName string) (*Kvs, error) {
	if dbName == "" {
		return nil, fmt.Errorf("empty database name is not valid")
	}
	if addr == "" {
		addr = ":1234"
	}
	go log.Fatal(http.ListenAndServe(addr, nil))
	return open(dbName)
}
