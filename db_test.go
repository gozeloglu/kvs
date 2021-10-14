package kvs

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tmpDb, err := New(t.Name())
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

func TestNewExistsFile(t *testing.T) {
	tmpDb, err := New(t.Name())
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

	tmpDb2, err := New(t.Name())
	if err == nil || tmpDb2 != nil {
		t.Fatalf("It should not be nil.")
	}
	t.Logf(err.Error())

	err = os.Remove(tmpDb.Dir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("TmpDb removed.")
}
