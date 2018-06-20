// Package systems is for system layor program
package systems

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// CreateText makes .txt file & writes byte-typed data in it
func CreateText() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File examples"))
	file.Close()
}

// StdOut sends example texts to stdout
func StdOut() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

// Buffer puts byte-typed data into buffer, and sends it to stdout
func Buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// FlushBuf uses buffer(write some texts) and flushes written data
func FlushBuf() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}

// MakeCopy reads an existing file and makes a copy
func MakeCopy() {
	oldFile, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create("goodbye.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)
}

// RandCreate makes a file filled with Rand data
func RandCreate() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)
}
