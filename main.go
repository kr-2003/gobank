package main

import (
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Abhinav", "Kumar", "hunterxhunter")
}

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	seedAccounts(store)
	server := NewAPIServer(":3000", store)
	server.Run()
	fmt.Println("Hello")
}
