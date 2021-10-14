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
	err = tmpDb.DbFile.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = os.Remove(tmpDb.Dir)
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
	err = tmpDb.DbFile.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb created: %s", t.Name())

	tmpDb2, err := Open(t.Name())
	if err != nil {
		t.Fatalf("It should not be nil.")
	}
	if tmpDb2.Name != t.Name() {
		t.Fatalf("Db name is wrong: %s", tmpDb2.Name)
	}
	tmpDb2.DbFile.Close()

	err = os.Remove(tmpDb.Dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb removed.")
}
