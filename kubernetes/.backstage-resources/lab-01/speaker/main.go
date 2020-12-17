package main

import (
	"fmt"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

func main() {
	fmt.Println("Hi there! I can't stop talking...")
	time.Sleep(500 * time.Millisecond)
	for {
		fmt.Printf("I say: %s\n", petname.Name())
		time.Sleep(500 * time.Millisecond)
	}
}
