package main

import (
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("sayHi", js.FuncOf(sayHi))
	js.Global().Set("sayHello", js.FuncOf(sayHello))
	js.Global().Set("callBack", js.FuncOf(callBack))
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("ex", js.FuncOf(ex))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
// GOARCH=wasm GOOS=js go build -o assets/lib.wasm cmd/lib/main.go
// cd cmd/lib
// GOARCH=wasm GOOS=js go build -o ../../assets/lib.wasm