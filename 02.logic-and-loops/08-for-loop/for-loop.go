package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "For normal")
	}

	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, "For array lenght")
	}

	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}
