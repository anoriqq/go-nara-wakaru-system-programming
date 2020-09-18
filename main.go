package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func copyN(dest io.Writer, src io.Reader, length int) (writeSize int64, err error) {
	reader := io.LimitReader(src, int64(length))
	return io.Copy(dest, reader)
}

func main() {
	reader := bytes.NewBufferString("exapmle of copyN\n")
	writeSize, err := io.CopyN(os.Stdout, reader, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(writeSize)
}
