package main

import (
	"fmt"

	"github.com/amonvix/golang/language_GO/concepts/1.base_concepts/02-fmt-package/poker"
)

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	for i := 1; i < 4; i++ {
		fmt.Printf("\nHand number: %v\n", i)
		poker.Deal()
	}

	// see emojis
	// Windows logo key + .
}
