// This is simple package, created in order to learn
// how to share module.
package greetings

import "fmt"

// Send a greeting message.
func Greet(name string) string {
	var mess string
	mess = fmt.Sprintf("Hi %v, thanks for using this package!", name)
	return mess
}
