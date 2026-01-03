package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	users []User `json:"users"`
}

type User struct {
	name       string  `json:"name"`
	occupation string  `json:"occupation"`
	age        int     `json:"age"`
	contact    Contact `json:"Contact"`
}

type Contact struct {
	email    string `json:"email"`
	whatsapp string `json:"whatsapp"`
}

func main() {
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()
	var users Users
	byteValue, _ := io.ReadAll((jsonFile))
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.users); i++ {
		fmt.Println("User name: " + users.users[i].name)
		fmt.Println("user age: " + strconv.Itoa(users.users[i].age))
	}
}
