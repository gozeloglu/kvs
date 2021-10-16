package kvs

import (
	"os"
	"testing"
	"time"
)

func TestOpen(t *testing.T) {
	tmpDb, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if _, err = os.Stat(baseDir + t.Name() + fileExtension); os.IsNotExist(err) {
		t.Fatalf("Directory was not created. %s", err.Error())
	}
	err = tmpDb.dbFile.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = os.Remove(tmpDb.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb removed.")
}

func TestOpenExistsFile(t *testing.T) {
	tmpDb, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if _, err = os.Stat(baseDir + t.Name() + fileExtension); os.IsNotExist(err) {
		t.Fatalf("Directory was not created. %s", err.Error())
	}
	err = tmpDb.dbFile.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb created: %s", t.Name())

	tmpDb2, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf("It should not be nil.")
	}
	if tmpDb2.name != t.Name() {
		t.Fatalf("Db name is wrong: %s", tmpDb2.name)
	}
	tmpDb2.dbFile.Close()

	err = os.Remove(tmpDb.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb removed.")
}

func TestKvs_Close(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")

	err = db.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB closed.")

	err = os.Remove(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("File removed.")
}

func TestWrite(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")

	err = db.load()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB loaded.")

	db.kv["a"] = "bar"
	db.kv["foo"] = "bar"

	err = db.write()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Data written to file.")
	db.Close()
	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf("Test file could not removed. %s", err.Error())
	}
	t.Logf("Removed test file.")
}

func TestLoad(t *testing.T) {
	db, err := open(t.Name(), "", 2*time.Minute)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB created.")
	db.kv["test"] = t.Name()

	err = db.write()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Data written to file.")

	err = db.load()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("DB loaded to memory.")

	if db.kv["test"] != t.Name() {
		t.Fatalf("wrong value. test=%s", db.kv["test"])
	}
	t.Logf("test=%s", db.kv["test"])

	err = os.RemoveAll(db.dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Removed test file.")

}
