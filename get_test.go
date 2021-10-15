package kvs

import (
	"os"
	"testing"
)

func TestKvs_GetEmpty(t *testing.T) {
	db, err := open(t.Name())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")

	key := "foo"

	v := db.Get(key)
	if v != "" {
		t.Fatalf("Value is wrong. Expected %s, got %s", "", v)
	}
	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Removed test directory.")
}

func TestKvs_Get(t *testing.T) {
	db, err := open(t.Name())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")

	key, value := "foo", "bar"
	db.Set(key, value)

	v := db.Get(key)
	if v != value {
		t.Fatalf("Value is wrong. Expected %s, got %s", value, v)
	}
	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Removed test directory.")
}
