package kvs

import (
	"fmt"
	"os"
)

type Kvs struct {
	// Name is database name.
	Name string

	// dir is directory of the data files.
	Dir string

	// DbFile is a File object that handles File operations.
	DbFile *os.File
}

const (
	baseDir       = "/tmp/kvs/"
	fileExtension = ".db"
)

// New creates data file for newly creating database. If the database file is
// already exists, it returns error without creating anything. name indicates
// database name.
func New(name string) (*Kvs, error) {
	fullPath := baseDir + name + fileExtension

	// Check database's base directory
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir(baseDir, 0700)
		if err != nil {
			return nil, fmt.Errorf("database directory couldn't created: %s", err.Error())
		}
	}

	// Check database file's directory
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		dbFile, err := os.Create(fullPath)
		if err != nil {
			return nil, fmt.Errorf("database couldn't created: %s", err.Error())
		}
		return &Kvs{
			Name:   name + fileExtension,
			Dir:    baseDir + name + fileExtension,
			DbFile: dbFile,
		}, nil
	}
	return nil, fmt.Errorf("this database already exists")
}
