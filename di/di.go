package di

import (
	"fmt"
	"io"
)

// Greet sends a personalized greeting to writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
