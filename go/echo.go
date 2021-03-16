package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	echo()
}

func echo() {
	for {
		r := bufio.NewReader(os.Stdin)
		s, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}
}

func echo2() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func echo3() {
	io.Copy(os.Stdout, os.Stdin)
}
