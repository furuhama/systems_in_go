// Package systems is for system layor program
package systems

import (
	"fmt"
	"os"
	"time"
)

// FormatTime send time to Stdout
func FormatTime() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
