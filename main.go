package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)
	// プロセスが落ちないのはなぜだろう
	io.Copy(os.Stdout, conn)
	req.Body.Close()
}
