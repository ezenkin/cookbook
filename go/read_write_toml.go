package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Name    string
	Address string
	Enabled bool
	ID      int
}

const filename = "test.toml"

func main() {
	writeFile(filename)
	readFile(filename)
	eraseFile(filename)
}

func writeFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Failed to open file", filename)
	}

	defer f.Close()

	conf := Config{"user", "127.0.0.1:554", true, 42}

	enc := toml.NewEncoder(f)
	enc.Encode(conf)
}

func readFile(filename string) {
	var conf Config
	if _, err := toml.DecodeFile(filename, &conf); err != nil {
		log.Fatalln("Failed to parse toml file", filename)
	}
	log.Println(conf)
}

func eraseFile(filename string) {
	os.Remove(filename)
}
