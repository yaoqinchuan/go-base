package base

import "fmt"

// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
func test() {
	defer func() {
		fmt.Println(recover()) //有效
	}()

	defer func() {
		panic("defer panic")
	}()
	defer recover() //无效
	defer fmt.Println(recover()) //无效

	defer func() {
		func() {
			fmt.Println("defer inner")
			fmt.Println(recover()) //无效
		}()
	}()

	panic("test panic")
}
/*
defer inner
<nil>
<nil>
defer panic

*/

// defer只在整个方法return或者抛出异常的时候才会逆向执行方法内的defer
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()

	fun()
}

// 使用defer和panic模拟java的try catch
func testTryCatch() {
	Try(func() {
		panic("this is a panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})

}
