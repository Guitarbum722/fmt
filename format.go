package strings

import (
	"fmt"
	"os"
)

// Printfln works exactly the same as the standard lib's Printf, but a new line is appended automatically
func Printfln(format string, a ...interface{}) (n int, err error) {
	format += "\n"
	return fmt.Fprintf(os.Stdout, format, a...)
}
