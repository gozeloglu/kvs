package kvs

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestServerSet(t *testing.T) {
	tmpFile, _ := os.CreateTemp(t.TempDir(), "setTest.kvs")
	t.Logf(tmpFile.Name())
	stat, err := tmpFile.Stat()
	if err != nil {
		t.Errorf(err.Error())
	}

	kvs := Kvs{
		name:     stat.Name(),
		dir:      tmpFile.Name(),
		dbFile:   tmpFile,
		kv:       make(map[string]string),
		mu:       sync.Mutex{},
		Addr:     "",
		duration: 10 * time.Minute,
	}

	body := `{"data": [{"key": "foo","value": "bar"}]}`
	//j, err := json.Marshal(body)
	//if err != nil {
	//	t.Fatalf(err.Error())
	//}
	req, err := http.NewRequest(http.MethodPost, "/set", strings.NewReader(body))
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(kvs.set)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"key":"","value":"","result":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	tmpFile.Close()
	err = os.RemoveAll(tmpFile.Name())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Removed temp directory and file.")
}

func TestServerGet(t *testing.T) {
	kvs := Kvs{
		name:     "",
		dir:      "",
		dbFile:   nil,
		kv:       nil,
		mu:       sync.Mutex{},
		Addr:     "",
		duration: 10 * time.Minute,
	}
	req, err := http.NewRequest(http.MethodGet, "/get/foo", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(kvs.get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"key":"foo","value":"","result":"OK"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	if rr.Header().Get(headerContent) != "application/json" {
		t.Errorf("handler returned wrong header content: got %s want %s",
			rr.Header().Get(headerContent), "application/json")
	}
}

func TestServerSave(t *testing.T) {
	tmpFile, _ := os.CreateTemp(t.TempDir(), "saveTest.kvs")
	t.Logf(tmpFile.Name())
	stat, err := tmpFile.Stat()
	if err != nil {
		t.Errorf(err.Error())
	}

	kvs := Kvs{
		name:     stat.Name(),
		dir:      tmpFile.Name(),
		dbFile:   tmpFile,
		kv:       make(map[string]string),
		mu:       sync.Mutex{},
		Addr:     "",
		duration: 10 * time.Minute,
	}
	req, err := http.NewRequest(http.MethodPut, "/save", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(kvs.save)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"key":"","value":"","result":"Saved"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	if rr.Header().Get(headerContent) != "application/json" {
		t.Errorf("handler returned wrong header content: got %s want %s",
			rr.Header().Get(headerContent), "application/json")
	}

	tmpFile.Close()
	err = os.RemoveAll(tmpFile.Name())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Removed temp directory and file.")
}
