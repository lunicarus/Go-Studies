package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name       string  `json:"name"`
	Occupation string  `json:"occupation"`
	Age        int     `json:"age"`
	Contact    Contact `json:"contacts"`
}

type Contact struct {
	Email    string `json:"email"`
	Whatsapp string `json:"whatsapp"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()
	var users Users
	byteValue, _ := io.ReadAll((jsonFile))
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User name: " + users.Users[i].Name)
		fmt.Println("user age: " + strconv.Itoa(users.Users[i].Age))
	}
}
