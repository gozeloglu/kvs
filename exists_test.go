package kvs

import (
	"os"
	"testing"
	"time"
)

func TestKvs_Exists(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)

	exist := db.Exists(k)
	if !exist {
		t.Errorf("memory does not contain the key: %s", k)
	}
	t.Logf("Exist value: %v", exist)

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("temp db removed.")
}

func TestKvs_NotExists(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)

	exist := db.Exists(t.Name())
	if exist {
		t.Errorf("unexpected result. key: %s", k)
	}
	t.Logf("Exist value: %v", exist)

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("temp db removed.")
}
