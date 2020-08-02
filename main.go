package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	writer := csv.NewWriter(multiWriter)
	writer.WriteAll([][]string{
		{"first_name", "last_name"},
		{"Rob", "Thompson"},
	})
}
