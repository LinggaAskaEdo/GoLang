package wrapper

import "fmt"

func wrap(f interface{}) interface{} {
	switch f2 := f.(type) {
	case func(i int) (ret int):
		return func(i int) (ret int) {
			fmt.Println("Before func(i int) (ret int), i =", i)
			ret = f2(i)
			fmt.Println("After func(i int) (ret int), ret =", ret)
			return
		}
	case func():
		return func() {
			fmt.Println("Before func()")
			f2()
			fmt.Println("After func()")
		}
	}
	return nil
}

func myfunc(i int) int {
	fmt.Println("myfunc called with", i)
	return i * 2
}

func myfunc2() {
	fmt.Println("myfunc2 called")
}

func WrapperTest() {
	wf := wrap(myfunc).(func(int) int)
	ret := wf(2)
	fmt.Println("Returned:", ret)

	wf2 := wrap(myfunc2).(func())
	wf2()
}
