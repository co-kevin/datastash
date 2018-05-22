package main

import (
	"testing"
)

type Person struct {
	Name  string
	PhoneNumber string
}

func TestConnectMongo(t *testing.T) {
	connectMongo("mongodb://localhost:27017")
	insertMongoDocument("foo", "person", Person{
		Name:  "Tom",
		PhoneNumber: "1662777",
	})
}
