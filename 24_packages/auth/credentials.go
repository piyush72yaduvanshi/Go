package auth

import "fmt"

// agar LoginWithCredentials function ko kisi aur package se call karna hai to uska naam capital letter se start hona chahiye

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
