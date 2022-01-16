package base

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type User struct {
	name  string
	email string
}

type Computer struct {
	User
	price int64
}

type Product struct {
	user User
	name string
}

func (u User) getInfo(name, email string) (info string) {
	u.name = name
	u.email = email
	return fmt.Sprintf("name is  %s  email is %s", name, email)
}

func (u *User) getSetInfo(name, email string) (info string) {
	u.name = name
	u.email = email
	return fmt.Sprintf("name is  %s  email is %s", name, email)
}

// 展示基础的方法使用
func testBaseMethod() {
	user := new(User)
	user.name = "张三"
	user.email = "zhansang@qq.com"
	fmt.Println(user.getInfo("李四", "lisi@qq.com"))
	fmt.Println(user)

	fmt.Println(user.getSetInfo("王五", "wangwu@qq.com"))

	fmt.Println(user)
}

/*
name is  李四  email is lisi@qq.com
&{张三 zhansang@qq.com}
name is  王五  email is wangwu@qq.com
&{王五 wangwu@qq.com}
*/

// 匿名字段可以通过上层直接访问，但是非匿名字段就不行
func testAnonymousField() {
	computer := new(Computer)
	computer.name = "张三"
	computer.email = "zhansang@qq.com"
	computer.price = 10000
	fmt.Println(computer)

	product := new(Product)
	product.name = "产品"
	product.user.email = "zhansang@qq.com"
	product.user.name = "张三"

	// 方法集只对匿名字段生效，在访问的时候也仅仅能够在匿名字段时候不用详细路径
	computer.getSetInfo("李四", "lisi@qq.com")
	fmt.Println(product)
}

/*
&{{张三 zhansang@qq.com} 10000}
&{{张三 zhansang@qq.com} 产品}
*/

// defer 碰上闭包,为什么全是4？因为在整个defer方法出现调用时候会拷贝他的入参，但是一旦出现闭包的场景，defer就会读取返回前的那个实时值
func testDefer() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}
/*
4
4
4
4
4
 */


//表达式本身就是方法的一个别名
//Golang 表达式 ：根据调用者不同，方法分为两种表现形式: instance.method(args...) 与<type>.func(instance, args...)
//method expression，使用对象类型来进行调用，对象本身需要作为参数传入，因此可以传指针或者对象本身，
//method value ，使用对象本身实例调用，他会复制 receiver对象本身，因此不能传指针。
func testExpression() {
	user := User{name: "张三", email: "zhangsang@qq.com"}
	expression := (*User).getSetInfo

	fmt.Println(expression(&user, "李四", "lisi@qq.com"))

	// 可以对复制类型的使用指针类型的方法族，反之不行
	expression1 := user.getSetInfo
	fmt.Println(expression1("王五", "wangwu@qq.com"))

	fmt.Println(user)
}

/*
name is  李四  email is lisi@qq.com
name is  王五  email is wangwu@qq.com
{王五 wangwu@qq.com}
*/

// 使用error的三种方式
//1、系统抛
func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能小于0")
	}
	return 3.14 * radius * radius
}
func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	getCircleArea(-1)
	fmt.Println("结束")
}

/*
半径不能小于0
*/
// 2、返回异常,使用异常对象
func getCircleAreaOrError(radius float32) (area float32, err error) {
	if radius < 0 {
		err = errors.New("半径不能小于0")
		return
	}
	area = 3.14 * radius * radius
	return
}

func tesErrorObject() {
	area, err := getCircleAreaOrError(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

/*
半径不能小于0
*/

// 3、自定义error,本质上是自己定义一个error对象与其可用的方法
type PathError struct {
	path       string
	option     string
	createTime string
	message    string
}

func (pathError PathError) Error() string {
	return fmt.Sprintf("path is %s, op is %s, creat time is %s, message is %s",
		pathError.path, pathError.option, pathError.createTime, pathError.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			option:     "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func testErrorObjectAndMethod() {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
/*
get path error, path is /Users/5lmh/Desktop/go/src/test.txt, op is read, creat time is 2022-01-08 23:39:26.7012008 +0800 CST m=+0.013961201,
message is open /Users/5lmh/Desktop/go/src/test.txt: The system cannot find the path specified.
 */