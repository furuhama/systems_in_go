// Package systems is for system layor program
package systems

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

// TCPConnect makes connection by tcp protocol & returns its result to stdout
func TCPConnect() {
	conn, err := net.Dial("tcp", "furuhama.github.io:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: furuhama.github.io\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter sample"))
}

// Handling stands up local server
func Handling() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

// HandleHTML is extended ver `Handling()`, which parses HTML string
func HandleHTML() {
	conn, err := net.Dial("tcp", "furuhama.github.io:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: furuhama.github.io\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// print header info
	fmt.Println(res.Header)
	// print body info
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
