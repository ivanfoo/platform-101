package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().Unix())

		suggestions := []string{
			"beans",
			"spaghetti",
			"peas",
			"ramen",
			"lentils",
			"soup",
			"burger",
		}
		index := rand.Intn(len(suggestions))
		suggestion := suggestions[index]

		fmt.Printf("Feeling hungry? You should try our delicious %s\n", suggestion)

		time.Sleep(1 * time.Second)
	}
}
