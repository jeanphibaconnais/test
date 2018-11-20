package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	alert := js.Global().Get("Alert")
	alert.Invoke("Alert : web assembly")
}