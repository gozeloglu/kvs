package kvs

import (
	"os"
	"testing"
	"time"
)

func TestConvStr(t *testing.T) {
	i := 12
	s := convStr(i)
	if s != "12" {
		t.Fatalf("String conversion is wrong.")
	}
	t.Logf("Converted successfully: %s", s)
}

func TestKvs_Incr(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "age", "12"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	val, err := db.Incr(k)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if val != "13" {
		t.Errorf("incremented value is wrong: %s", val)
	}
	t.Logf("value is incremented.")
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

func TestKvs_IncrStr(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	_, err = db.Incr(k)
	if err == nil {
		t.Fatalf("error expected, but nil got.")
	}
	t.Logf("value could not incremented. it is string.")

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

func TestKvs_IncrBy(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v, i := "age", "12", 3
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	val, err := db.IncrBy(k, i)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if val != "15" {
		t.Errorf("incremented value is wrong: %s", val)
	}
	t.Logf("value is incremented.")
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

func TestKvs_IncrByStr(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("db created.")

	k, v := "foo", "bar"
	db.Set(k, v)
	t.Logf("Key-value pair is set.")
	t.Logf("%s: %s", k, v)
	_, err = db.IncrBy(k, 3)
	if err == nil {
		t.Fatalf("error expected, but nil got.")
	}
	t.Logf("value could not incremented. it is string.")

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
