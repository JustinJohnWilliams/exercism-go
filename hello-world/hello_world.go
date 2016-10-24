package hello

import "fmt"

const testVersion = 2

// HelloWorld will print "Hello, [input]" or "Hello, World" if no input received.
func HelloWorld(s string) string {

	if s != "" {
		return fmt.Sprintf("Hello, %s!", s)
	}
	return "Hello, World!"
}
