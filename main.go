package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("formated.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "No.%d %s weights %f kg\n", 1, "Bob", 65.3)
	file.Close()
}
