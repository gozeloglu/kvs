package kvs

import (
	"fmt"
	"os"
)

type Kvs struct {
	// name is database name.
	name string

	// dir is directory of the data files.
	dir string

	// dbFile is a File object that handles File operations.
	dbFile *os.File
}

const (
	baseDir       = "/tmp/kvs/"
	fileExtension = ".kvs"
)

// Open creates data file for newly creating database. If the database file is
// already exists, it returns error without creating anything. name indicates
// database name.
func Open(name string) (*Kvs, error) {
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
			name:   name + fileExtension,
			dir:    baseDir + name + fileExtension,
			dbFile: dbFile,
		}, nil
	} else {
		return open(name)
	}
}

// open opens the named database for file operations.
func open(dbName string) (*Kvs, error) {
	fullPath := baseDir + dbName + fileExtension
	dbFile, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	return &Kvs{
		name:   dbName,
		dir:    fullPath,
		dbFile: dbFile,
	}, nil
}

// Close closes the file.
func (k *Kvs) Close() error {
	return k.dbFile.Close()
}
