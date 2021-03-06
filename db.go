package kvs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// Kvs keeps the essential variables.
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

	// Addr is the server address.
	Addr string

	// duration stands for the time interval to save data into file periodically.
	duration time.Duration
}

const (
	baseDir       = "/tmp/kvs/"
	fileExtension = ".kvs"
)

// open creates data file for newly creating database. If the database file is
// already exists, it returns error without creating anything. name indicates
// database name.
func open(name string, addr string, duration time.Duration) (*Kvs, error) {
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
		k := &Kvs{
			name:     name + fileExtension,
			dir:      baseDir + name + fileExtension,
			dbFile:   dbFile,
			kv:       m,
			mu:       sync.Mutex{},
			Addr:     addr,
			duration: duration,
		}

		ticker := time.NewTicker(duration)
		go func() {
			for {
				select {
				case t := <-ticker.C:
					err := k.write()
					if err != nil {
						log.Println("Writing file failed at", t.Local())
					} else {
						log.Println("Data saved on the file at", t.Local())
					}
				}
			}
		}()
		return k, nil
	}
	k, err := openAndLoad(name, addr, duration)
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				err := k.write()
				if err != nil {
					log.Println("Writing file failed at", t.Local())
				} else {
					log.Println("Data saved on the file at", t.Local())
				}
			}
		}
	}()
	return k, nil
}

// openAndLoad opens the named database file for file operations. Also, loads
// the database file into map to in-memory operations.
func openAndLoad(dbName string, addr string, duration time.Duration) (*Kvs, error) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	fullPath := baseDir + dbName + fileExtension
	dbFile, err := os.OpenFile(fullPath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	k := &Kvs{
		name:     dbName,
		dir:      fullPath,
		dbFile:   dbFile,
		mu:       sync.Mutex{},
		Addr:     addr,
		duration: duration,
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
	defer k.dbFile.Close()
	d := ""
	for key, val := range k.kv {
		d += fmt.Sprintf("%s=%s\n", key, val)
	}
	return ioutil.WriteFile(k.dir, []byte(d), 0666)
}
