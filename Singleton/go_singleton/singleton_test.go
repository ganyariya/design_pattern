package go_singleton

import (
	"testing"
)

func getInstance(database *Database) {
	database = getDatabaseInstance()
}

func TestSingleton(t *testing.T) {
	var database *Database = nil
	for i := 0; i < 30; i++ {
		go getInstance(database)
	}

	if database.getName() != "Database" {
		t.Fatalf("Error")
	}
}
