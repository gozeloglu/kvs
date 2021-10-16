package kvs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type Kvs struct {
	// name is database name.
	name string

	// dir is directory of the data files.
	dir string

	// dbFile is a File object that handles File operations.
	dbFile *os.File

	// kv means key-value. It keeps the data in map.
	kv map[string]string

	// mu is mutex for avoiding conflicts.
	mu sync.Mutex

	// addr is the server address.
	addr string
}

const (
	baseDir       = "/tmp/kvs/"
	fileExtension = ".kvs"
)

// open creates data file for newly creating database. If the database file is
// already exists, it returns error without creating anything. name indicates
// database name.
func open(name string, addr string) (*Kvs, error) {
	fullPath := baseDir + name + fileExtension

	// Check database's base directory
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir(baseDir, 0777)
		if err != nil {
			return nil, fmt.Errorf("database directory couldn't created: %s", err.Error())
		}
	}

	// Check database file's directory
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		dbFile, err := os.OpenFile(fullPath, os.O_CREATE, 0777)
		if err != nil {
			return nil, fmt.Errorf("database couldn't created: %s", err.Error())
		}
		m := make(map[string]string)
		return &Kvs{
			name:   name + fileExtension,
			dir:    baseDir + name + fileExtension,
			dbFile: dbFile,
			kv:     m,
			mu:     sync.Mutex{},
			addr:   addr,
		}, nil
	} else {
		return openAndLoad(name, addr)
	}
}

// openAndLoad opens the named database file for file operations. Also, loads
// the database file into map to in-memory operations.
func openAndLoad(dbName string, addr string) (*Kvs, error) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	fullPath := baseDir + dbName + fileExtension
	dbFile, err := os.OpenFile(fullPath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	k := &Kvs{
		name:   dbName,
		dir:    fullPath,
		dbFile: dbFile,
		mu:     sync.Mutex{},
		addr:   addr,
	}
	err = k.load()
	if err != nil {
		return nil, err
	}

	return k, nil
}

// Close closes the file.
func (k *Kvs) Close() error {
	return k.write()
}

// load reads and loads the data from the file into map.
func (k *Kvs) load() error {
	k.mu.Lock()
	defer k.mu.Unlock()
	m := make(map[string]string)
	buf, err := os.ReadFile(k.dir)

	if err != nil {
		return err
	}

	fileData := string(buf[:])
	if fileData == "" {
		k.kv = m
		return nil
	}
	dataArr := strings.Split(fileData, "\n")
	for i := 0; i < len(dataArr)-1; i++ {
		data := strings.Split(dataArr[i], "=")
		k, v := data[0], data[1]
		m[k] = v
	}
	k.kv = m
	return nil
}

// write saves data into file. It writes the data in map to the file.
func (k *Kvs) write() error {
	k.mu.Lock()
	defer k.mu.Unlock()
	d := ""
	for key, val := range k.kv {
		d += fmt.Sprintf("%s=%s\n", key, val)
	}
	err := k.dbFile.Close()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(k.dir, []byte(d), 0666)
}
