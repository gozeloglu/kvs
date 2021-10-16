package kvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type KeyValue struct {
	Data []Data `json:"data"`
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Result string `json:"result"`
}

// Create creates database and Kvs object. It creates database and returns Kvs
// object. If HTTP address is empty, localhost and default port is used.
// In contrast, dbName name needs to be specified. If it is not specified, it
// returns error and the database is not created.
func Create(addr string, dbName string) (*Kvs, error) {
	if dbName == "" {
		return nil, fmt.Errorf("empty database name is not valid")
	}
	if addr == "" {
		addr = "localhost:1234"
	}
	return open(dbName, addr)
}

// Open creates an HTTP connection. HTTP connection listens HTTP  requests from
// client. Create function needs to be called before calling Open function.
func (k *Kvs) Open() {
	log.Printf("Kvs server running on %s...", k.addr)
	http.HandleFunc("/set", k.set)
	http.HandleFunc("/get", get)
	log.Fatal(http.ListenAndServe(k.addr, nil))
}

// set is the /set API endpoint for setting a key-value pair.
func (k *Kvs) set(w http.ResponseWriter, r *http.Request) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if r.Method != http.MethodPost {
		err := fmt.Sprintf("Wrong HTTP request. You need to send POST request.")
		log.Printf(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var keyVal *KeyValue
	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(keyVal.Data); i++ {
		key, val := keyVal.Data[i].Key, keyVal.Data[i].Value
		k.kv[key] = val
	}
	data := Response{
		Result: "OK",
	}
	j, err := json.Marshal(data)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	log.Printf("Key-value pair is set.")
}

func get(w http.ResponseWriter, r *http.Request) {
	// TODO Handle get operation.
}
