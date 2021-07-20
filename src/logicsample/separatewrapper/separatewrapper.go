package separatewrapper

import "fmt"

func trace(funcName string) func() {
	fmt.Println("START method", funcName)
	return func() {
		fmt.Println("END method", funcName)
	}
}

func doSomething(name string) {
	defer trace("something")()
	fmt.Println(name)
}

func SeparateWrapperTest() {
	doSomething("test")
}
