package kvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

const (
	headerContent = "Content-Type"
	contentValue  = "application/json"
)

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
	http.HandleFunc("/get/", k.get)
	http.HandleFunc("/save", k.save)
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
	w.Header().Set(headerContent, contentValue)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	log.Printf("Key-value pair is set.")
}

// get returns the value of the key.
func (k *Kvs) get(w http.ResponseWriter, r *http.Request) {
	k.mu.Lock()
	defer k.mu.Unlock()

	if r.Method != http.MethodGet {
		err := fmt.Sprintf("Wrong HTTP request. You need to send GET request.")
		log.Printf(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	u := strings.Split(r.URL.String(), "/")
	if u[len(u)-1] == "" && u[len(u)-2] == "get" {
		err := fmt.Sprintf("Key is missing.")
		log.Printf(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}
	key := u[len(u)-1]
	value := k.Get(key)

	resp := Response{
		Key:    key,
		Value:  value,
		Result: "OK",
	}
	j, err := json.Marshal(resp)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(headerContent, contentValue)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	log.Printf("%s=%s", key, value)
}

func (k *Kvs) save(w http.ResponseWriter, r *http.Request) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if r.Method != http.MethodPut {
		err := fmt.Sprintf("Wrong HTTP request. You need to send PUT request.")
		log.Printf(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	err := k.write()
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := Response{
		Result: "Saved",
	}
	j, err := json.Marshal(resp)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(headerContent, contentValue)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	log.Printf("Saved.")
}
