package kvs

import (
	"os"
	"testing"
)

func TestKvs_Set(t *testing.T) {
	db, err := open(t.Name())
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
