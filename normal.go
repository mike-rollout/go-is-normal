package normal

import (
	"fmt"
	"math/rand"
	"time"
)

// isNormal randomly returns "yes" or "no" and prints a funny motivational quote
func IsNormal() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		fmt.Println("You are as normal as a unicorn dancing on a rainbow!")
		return "no"
	} else {
		fmt.Println("Normal is just a setting on the washing machine!")
		return "yes"
	}
}
