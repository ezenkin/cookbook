package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
