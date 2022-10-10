package go_singleton

import (
	"fmt"
	"sync"
)

type Database struct{}

func (*Database) getName() string {
	return "Database"
}

var lock = &sync.Mutex{}
var database *Database

func getDatabaseInstance() *Database {
	if database == nil {
		lock.Lock()
		defer lock.Unlock()
		if database == nil {
			fmt.Println("Create instance now...")
			database = &Database{}
		} else {
			fmt.Println("Already created...")
		}
	}
	return database
}
