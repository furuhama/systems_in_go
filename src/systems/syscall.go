// Package systems is for system layor program
package systems

import (
	"os"
)

// FetchOSCreate fetch os.Create
func FetchOSCreate() {
	// trace definition od `os.Create()`
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("system call example"))
}

// [tracing path]
// -> /os/file.go#246 func Create
// -> /os/file_unix.go#147 func OpenFile()
// -> /syscall/zsyscall_darwin_amd64.go#881 func Open
// -> (/syscall/syscall_unix.go#29 func Syscall)
// -> /syscall/asm_darwin_amd64.s#15 TEXT> ·Syscall(SB),NOSPLIT,$0-56
