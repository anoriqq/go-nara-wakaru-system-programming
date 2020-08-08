package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	sourceFileName := flag.String("s", "old.txt", "source file name")
	outputFileName := flag.String("o", "new.txt", "output file name")
	flag.Parse()

	file, err := os.Open(*sourceFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create(*outputFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, file)
}
