package main

import "fmt"

func main() {
	users := getUsers()

	for _, user := range users {
		fmt.Println("User Id: ", user.id, "User Name: ", user.name)
	}

	P := struct {
		name string
		age  int
	}{"Amy", 18} // Anonnymous struct

	fmt.Println(P)
}

func getUsers() []user {
	users := []user{
		user{
			id:       1,
			name:     "User 01",
			password: "Senha 123",
		},
		user{
			id:       2,
			name:     "User 02",
			password: "Senha 312",
		},
	}

	return users
}

type id int
type user struct {
	id       id
	name     string
	password string
}
