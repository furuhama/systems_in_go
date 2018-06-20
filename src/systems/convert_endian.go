// Package systems is for system layor program
package systems

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ConvertEndian converts big endian to little endian
func ConvertEndian() {
	// data is 32bit big endian value (which means 10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// convert
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
