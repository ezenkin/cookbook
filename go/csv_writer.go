package main

import (
	"encoding/csv"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)

	records := [][]string{
		{"date", "one", "two", "three"},
		{"05/11/2018", "1", "2", "3"},
	}

	w.Write(records[0])
	w.Flush()

	w.WriteAll(records)
	w.Flush()
}
