package kvs

import (
	"os"
	"testing"
	"time"
)

func TestKvs_Decr(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "age", "12"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	val, err := db.Decr(k)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if val != "11" {
		t.Errorf("decremented value is wrong: %s", val)
	}
	t.Logf("value is decremented.")
	t.Logf("%s: %s", k, val)

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

func TestKvs_DecrStr(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	_, err = db.Decr(k)
	if err == nil {
		t.Fatalf("error expected, but nil got.")
	}
	t.Logf("value could not decremented. it is string.")

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
