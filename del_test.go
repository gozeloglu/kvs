package kvs

import (
	"os"
	"testing"
	"time"
)

func TestDel(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)

	val := db.Get(k)
	if val != v {
		t.Logf("set value could not get.")
	}

	db.Del(k)
	if db.Get(k) != "" {
		t.Logf("key could not deleted.")
	}

	t.Logf("Key is deleted successfully.")

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("test file removed.")
}

func TestDelWithTime(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Second)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	val := db.Get(k)
	t.Logf("%s: %s", k, val)
	db.Del(k)
	val = db.Get(k)
	if val != "" {
		t.Errorf("val is not empty string: %s", val)
	}
	time.Sleep(3 * time.Second)

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db closed.")

	tmpDb, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("tmpdb created.")

	val = tmpDb.Get(k)
	if val != "" {
		t.Errorf("val is not empty string: %s", val)
	}

	err = tmpDb.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db closed.")

	err = os.RemoveAll(tmpDb.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("tmp file removed.")
}
