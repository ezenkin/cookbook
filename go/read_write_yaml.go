package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type User struct {
	Name      string
	Age       int
	Birthdate time.Time
}

type Config []User

const fileName = "config.yml"

const data = `
- name: Tom
  age: 42
  birthdate: 1979-02-25
- name: Jain
  age: 30
  birthdate: 1991-02-25
`

func main() {
	/*
		cfg := Config{
			User{
				Name:      "Tom",
				Age:       42,
				Birthdate: time.Now().AddDate(-42, 0, 0),
			},
			User{
				Name:      "Jain",
				Age:       30,
				Birthdate: time.Now().AddDate(-30, 0, 0),
			},
		}
	*/

	var cfg Config

	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		log.Fatalln(err)
	}
	log.Println("Write config to file:", fileName)
	log.Println(cfg)

	if err := cfg.Write(fileName); err != nil {
		log.Fatalln(err)
	}

	var newCfg Config
	if err := newCfg.Read(fileName); err != nil {
		log.Fatalln(err)
	}

	log.Println("Read config from file:", fileName)
	log.Println(newCfg)

	log.Println("Erase config file:", fileName)
	if err := os.Remove(fileName); err != nil {
		log.Fatalln(err)
	}
}

func (cfg *Config) Write(filename string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	defer enc.Close()
	return enc.Encode(cfg)
}

func (cfg *Config) Read(filename string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	return dec.Decode(cfg)
}
