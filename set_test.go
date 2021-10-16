package kvs

import (
	"os"
	"testing"
	"time"
)

func TestKvs_Set(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")

	key, value := "foo", "bar"
	db.Set(key, value)

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB closed.")

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Test file removed.")
}
