package kvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// KeyValue is used for unmarshalling JSON object in POST request.
type KeyValue struct {
	Data []Data `json:"data"`
}

// Data is the element of array. It keeps Key and Value.
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Response is struct type for JSON response body. It is used for get and save.
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
// returns error and the database is not created. Kvs saves the data inside the
// map to the file periodically. User needs to specify the time interval as
// duration. For example, if 2*time.Minute is passed to duration parameter,
// the data that stores in memory, map, saves to the file.
func Create(addr string, dbName string, duration time.Duration) (*Kvs, error) {
	if dbName == "" {
		return nil, fmt.Errorf("empty database name is not valid")
	}
	if addr == "" {
		addr = "localhost:1234"
	}
	return open(dbName, addr, duration)
}

// Open creates an HTTP connection. HTTP connection listens HTTP  requests from
// client. Create function needs to be called before calling Open function.
func (k *Kvs) Open() {
	log.Printf("Kvs server running on %s...", k.Addr)
	http.HandleFunc("/set", k.set)
	http.HandleFunc("/get/", k.get)
	http.HandleFunc("/save", k.save)
	log.Fatal(http.ListenAndServe(k.Addr, nil))
}

// set is the /set API endpoint for setting a key-value pair.
func (k *Kvs) set(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP method: %s", r.Method)
	log.Printf("Endpoint: %s", r.URL)
	log.Printf("Request header: %s", r.Header)
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
	log.Printf("Body read.")

	var keyVal *KeyValue
	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Body unmarshalled to KeyVal struct type.")

	for i := 0; i < len(keyVal.Data); i++ {
		key, val := keyVal.Data[i].Key, keyVal.Data[i].Value
		k.Set(key, val)
	}
	log.Printf("Save key-value pair to memory.")

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
	_, err = w.Write(j)
	if err != nil {
		return
	}
	log.Printf("Status code: %d", http.StatusOK)
	log.Printf("Response header: %s", w.Header().Get(headerContent))
}

// get returns the value of the key.
func (k *Kvs) get(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP Method: %s", r.Method)
	log.Printf("Endpoint: %s", r.URL)
	log.Printf("Request Header: %s", r.Header)
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
	log.Printf("URL parsed.")

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
	log.Printf("Response body marshalled.")

	w.Header().Set(headerContent, contentValue)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		return
	}
	log.Printf("%s=%s", key, value)
	log.Printf("Status code: %d", http.StatusOK)
	log.Printf("Response header: %s", w.Header().Get(headerContent))
}

// save writes the data from map to file.
func (k *Kvs) save(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP method: %s", r.Method)
	log.Printf("Endpoint: %s", r.URL)
	log.Printf("Request header: %s", r.Header)
	if r.Method != http.MethodPut {
		err := fmt.Sprintf("Wrong HTTP request. You need to send PUT request.")
		log.Printf(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	k.mu.Lock()
	err := k.write()
	k.mu.Unlock()
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Data saved to database file.")

	resp := Response{
		Result: "Saved",
	}
	j, err := json.Marshal(resp)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Response body marshalled.")

	w.Header().Set(headerContent, contentValue)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		return
	}
	log.Printf("Status code: %d", http.StatusOK)
	log.Printf("Response header: %s", w.Header().Get(headerContent))
}
