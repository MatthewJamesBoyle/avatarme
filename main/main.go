package main

import (
	"flag"
	"AvatarGenerator"
)

func main() {

	email := flag.String("email", "", "email address of the user")
	ip := flag.String("ip", "", "ip address of the user")
	user := flag.String("username", "", "username of user")
	flag.Parse()

	AvatarGenerator.GenerateAvatar(*email, *ip, *user)

}
