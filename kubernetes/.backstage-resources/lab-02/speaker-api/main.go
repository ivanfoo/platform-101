package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	petname "github.com/dustinkirkland/golang-petname"
)

var totalWords int64

func main() {
	fmt.Println("Hi there! I can't stop talking...")
	go talk()

	r := gin.Default()
	r.GET("/words", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"total": totalWords,
		})
	})
	r.Run()
}

func talk() {
	for {
		fmt.Printf("I say: %s\n", petname.Name())
		totalWords = totalWords + 1
		time.Sleep(500 * time.Millisecond)
	}
}
