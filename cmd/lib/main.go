package main
import (
	"fmt"
	"syscall/js"
)

func sayHi(value js.Value, args []js.Value) interface{} {
	fmt.Println("Hi!")
	fmt.Printf("value:%v\n", value)
	fmt.Printf("args:%v\n", args)
	return nil
}

func sayHello(value js.Value, args []js.Value) interface{} {
	fmt.Println("sayHello!")
	fmt.Printf("value:%v\n", value)
	fmt.Printf("args:%v\n", args)
	callback := args[0].String()
	f := js.Global().Get(callback)
	message := "sssss"
	// 调用函数
	f.Invoke(message)
	//callback.Invoke(js.Null(), "Did you say "+message)
	return nil
}

func registerCallbacks() {
	js.Global().Set("sayHi", js.FuncOf(sayHi))
	js.Global().Set("sayHello", js.FuncOf(sayHello))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
// GOARCH=wasm GOOS=js go build -o assets/lib.wasm cmd/lib/main.go