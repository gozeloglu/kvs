package kvs

import (
	"os"
	"testing"
	"time"
)

func TestKvs_Rename(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)

	ok, err := db.Rename(k, k+k)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !ok {
		t.Errorf("key didn't change")
	} else {
		t.Logf("key renamed.")
	}

	val := db.Get(k)
	if val != "" {
		t.Errorf("%s-%s pair is still exits.", k, val)
	}
	t.Logf("%s-%s", k, val)

	val = db.Get(k + k)
	if val != v {
		t.Errorf("expected %s-%s pair, but got %s-%s", k+k, v, k+k, val)
	}
	t.Logf("%s-%s", k+k, val)

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

func TestKvs_RenameExistsNewKey(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)

	ok, err := db.Rename(k, k)
	if !ok {
		t.Logf("key didn't change")
	}

	val := db.Get(k)
	if val != v {
		t.Errorf("%s-%s pair is wrong.", k, val)
	}

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
