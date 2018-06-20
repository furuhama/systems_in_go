// Package systems is for system layor program
package systems

import (
	"fmt"
	"os"
	"time"
)

// FormatTime sends time to stdout
func FormatTime() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
