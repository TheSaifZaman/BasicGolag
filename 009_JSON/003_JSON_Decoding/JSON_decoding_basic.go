package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Password string
}

func main() {
	var jsoninput = []byte("[{\"Username\":\"debora\",\"Password\":\"123456\"},{\"Username\":\"bob\",\"Password\":\"42\"},{\"Username\":\"sandra\",\"Password\":\"33\"}]")
	var users []user
	err := json.Unmarshal(jsoninput, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	for _, user := range users {
		fmt.Println("->", user.Username)
	}
}
