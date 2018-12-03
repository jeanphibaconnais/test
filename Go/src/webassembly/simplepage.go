package main

import (
	"syscall/js"
	"fmt"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	alert := js.Global().Get("Alert")
	alert.Invoke("Alert : web assembly")
}