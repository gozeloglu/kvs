package kvs

import (
	"os"
	"testing"
	"time"
)

func TestKvs_Keys(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	pairs := []string{"foo", "1", "bar", "2", "fizz", "3", "buzz", "4"}
	for i := 0; i < len(pairs); i += 2 {
		db.Set(pairs[i], pairs[i+1])
	}
	t.Logf("key-value pairs are set.")

	keys := db.Keys()
	if len(keys) != len(pairs)/2 {
		t.Errorf("Fetched keys are wrong. pairs len is %d, "+
			"fetched len is %d.", len(pairs)/2, len(keys))
	}
	for _, k := range keys {
		t.Logf(k)
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

func TestKvs_EmptyKeys(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	keys := db.Keys()
	if len(keys) != 0 {
		t.Errorf("Fetched keys are wrong. pairs len is %d, "+
			"fetched len is %d.", 0, len(keys))
	}
	t.Logf("Empty slice: %d", len(keys))

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
