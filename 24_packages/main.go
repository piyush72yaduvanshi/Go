package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/podcast/auth"
	"github.com/podcast/user"
)

func main() {
	auth.LoginWithCredentials("piyush", "#@12345")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)
	color.Red(user.Name)

}

// go mod init github.com/podcast
//  iss command se go.mod file create hoti hai jisme github url mention hota hai jisse hum apne package ko import kar sakein

//  go package ke liye hme https://pkg.go.dev/   yahan jana hota hn  u can also search go package as npm package

// installation of package process
// 1. go get github.com/fatih/color  -- iss command se go.mod file update hoti hai
// 2. go mod tidy --  iss command se go.sum file create hoti hai jisme package ke dependencies ka details hota hai
// 3. go run main.go  -- iss command se program run hoti hai

//go.mod file -- package ke details hoti hai
//go.sum file -- package ke dependencies ka details hoti hai
