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

// CreateText is making txt file & write byte data in it
func CreateText() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File examples"))
	file.Close()
}

// StdOut send Stdout to example texts
func StdOut() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

// Buffer put byte data into buffer, and send it to Stdout
func Buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// FlushBuf uses Buffer and flush its data
func FlushBuf() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}

// MakeCopy reads exist file and make a copy
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
