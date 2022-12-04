// run with `go run ./demo/pkg/main`
package main

import (
	"lectures/demo/pkg/display"
	"lectures/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("calling a package function")
}
