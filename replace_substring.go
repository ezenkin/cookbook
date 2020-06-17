package main

import (
	"log"
	"strings"
)

func main() {
	str := "Hello, world, again."
	old, rep := "world", "earth"
	log.Println(str) //Hello, world, again.
	str = strings.Replace(str, old, rep, -1)
	log.Println(str) //Hello, earth, again.
}
