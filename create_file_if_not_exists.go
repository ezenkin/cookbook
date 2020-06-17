package main

import (
	"log"
	"os"
)

const fileName = "test.txt"

func main() {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Can't create file:", fileName)
		}
		log.Println("Create file:", fileName)
		f.Close()

		//due to program is only to test - remove file the testing is done
		log.Println("Remove file:", fileName)
		os.Remove(fileName)
	}
}
