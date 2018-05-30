package main

import (
	"testing"
	"time"
)

type Person struct {
	Name        string
	PhoneNumber string
	Date        time.Time
}

func TestConnectMongo(t *testing.T) {
	connectMongo("mongodb://localhost:27017")
	insertMongoDocument("foo", "person", Person{
		Name:        "Tom",
		PhoneNumber: "1662777",
		Date:        time.Now().Local(),
	})
}
