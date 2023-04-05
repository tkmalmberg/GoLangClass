package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	peopleSlice := []string{"Tyler", "Megan", "Ghibli", "Puddles", "Peppermint", "Peaches"}

	if x := peopleSlice[rand.Intn(len(peopleSlice))]; x == "Tyler" {
		fmt.Println("Tyler is working at the computer.")
	} else if x == "Megan" {
		fmt.Println("Megan is crochetting.")
	} else {
		fmt.Println("Ghibli is stinky.")
	}
}