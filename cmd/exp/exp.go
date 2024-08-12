package main

import (
	"errors"
	"fmt"
	"log"
)

type Reviews struct {
	Comment string
	Rating  int
}

type User struct {
	Name           string
	Bio            string
	Email          string
	FavouriteFoods map[string]string
	Reviews        []Reviews
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}

func main() {
	err := CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}
}
