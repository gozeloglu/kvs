package kvs

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	tmpDb, err := Open(t.Name())
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
	tmpDb, err := Open(t.Name())
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

	tmpDb2, err := Open(t.Name())
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
