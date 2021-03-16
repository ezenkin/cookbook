package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.LookupEnv("PATH"))
	fmt.Println()
	fmt.Println(os.Getenv("PATH"))
	fmt.Println()
	fmt.Println(os.Environ())
}
