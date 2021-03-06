package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	db2 "github.com/upper/db/v4"
)

func TestNew(t *testing.T) {

	fakeDB, _, _ := sqlmock.New()
	defer fakeDB.Close()
	//postgres
	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("wrong type:", fmt.Sprintf("%T", m))
	}

	// mysql
	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m = New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("wring type:", fmt.Sprintf("%T", m))
	}
}

func TestGetInsertID(t *testing.T) {
	var id db2.ID
	id = int64(1)

	returnedID := getInsertID(id)

	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("wrong type")
	}

	id = 1

	returnedID = getInsertID(id)

	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("wrong type")
	}
}
