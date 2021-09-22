package main

import (
	"fmt"
	"os/exec"
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

func callBack(value js.Value, args []js.Value) interface{} {
	fmt.Println("callback!")
	fmt.Printf("value:%v\n", value)
	fmt.Printf("args:%v\n", args)
	// 函数名字
	name := args[0].String()
	f := js.Global().Get(name)
	// 调用函数
	f.Invoke(args[1])
	return nil
}

func add(value js.Value, args []js.Value) interface{}  {
	if len(args) != 2 {
		return nil
	}
	js.Global().Set("output", js.ValueOf(args[0].Int()+args[1].Int()))
	fmt.Printf("ret:%v\n", js.ValueOf(args[0].Int() + args[1].Int()).String())
	return nil
}

func subtract(value js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}
	js.Global().Set("output", js.ValueOf(args[0].Int()-args[1].Int()))
	fmt.Printf("ret:%v\n", js.ValueOf(args[0].Int() -args[1].Int()).String())
	return nil
}

func ex(value js.Value, args []js.Value) interface{} {
	fmt.Println("ex!")
	fmt.Printf("value:%v\n", value)
	fmt.Printf("args:%v\n", args)
	cmdName := ""
	argsStr := make([]string, 0)
	for idx, arg := range args {
		if idx == 0 {
			cmdName = arg.String()
			continue
		}
		argsStr = append(argsStr, arg.String())
	}
	cmd := exec.Command(cmdName, argsStr...)
	//err := cmd.Run()
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("err:%v", err)
		return nil
	}
	fmt.Printf("out:%v", string(out))
	return nil
}