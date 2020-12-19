package main

import "fmt"

func main() {
	users := getUsers()
	printUsers(users)
	printSeparator()
	addUser(users, "João", "Souza")
	printUsers(users)
	printSeparator()
	deleteUser(users, "João")
	printUsers(users)
}

func printSeparator() {
	fmt.Println("==========================")
}

func getUsers() map[string]string {
	return map[string]string{
		"Carol":  "Monserrat",
		"Filipe": "Souza",
	}
}

func addUser(users map[string]string, key string, value string) {
	users[key] = value
}

func printUsers(users map[string]string) {
	for key, value := range users {
		fmt.Println("  ID:", key, "Name:", value)
	}
}

func deleteUser(users map[string]string, key string) {
	delete(users, key)
}
