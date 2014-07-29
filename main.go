package main

import (
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	fmt.Printf("%x\n", securecookie.GenerateRandomKey(64))
	fmt.Printf("%x\n", securecookie.GenerateRandomKey(32))
}
