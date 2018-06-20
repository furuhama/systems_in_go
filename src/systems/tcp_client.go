// Package systems is for system layor program
package systems

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
)

// TCPSocketClient sends http request
// to be compatible with keep-alive
func TCPSocketClient() {
	sendMessages := []string{
		"PIYO",
		"FUNCTION",
		"PLUSPLUS",
	}
	current := 0
	var conn net.Conn
	for {
		var err error
		// start Dial when haven't set connection, or retry because of an error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// define request to send literals as POST method
		request, err := http.NewRequest("POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		request.Header.Set("Accept-Encoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// read response from server
		// when timeout, an error occurs written below
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// show results
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		// end if every content is transported
		current++
		if current == len(sendMessages) {
			break
		}
	}
}

// TCPSocketClientChunk sends http request
// to get chunked data
func TCPSocketClientChunk() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	// define request as GET method
	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	// read response from server
	// when timeout, an error occurs written below
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	// show results
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		// get size of transferred data
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		// parse hex size
		// close if size if zero
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		// read data by getting buffer as size length
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}
}

// TCPSocketClientPipeline sends http request
// to get data with pipeline
func TCPSocketClientPipeline() {
	sendMessages := []string{
		"PIYO",
		"HOGEFUGA",
		"FUNCTIONAL",
	}
	current := 0
	var conn net.Conn
	var err error
	requests := make(chan *http.Request, len(sendMessages))
	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access: %d\n", current)
	defer conn.Close()

	// send just request at first
	for i := 0; i < len(sendMessages); i++ {
		isLastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest("GET",
			"http://localhost:8888?message="+sendMessages[i],
			nil)
		if isLastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-aive")
		}
		if err != nil {
			panic(err)
		}

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		fmt.Println("send: ", sendMessages[i])
		requests <- request
	}
	close(requests)

	// get all responses at once
	reader := bufio.NewReader(conn)
	for request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		if current == len(sendMessages) {
			break
		}
	}
}
