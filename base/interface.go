package base

/*
Go语言提倡面向接口编程。
接口是一个或多个方法签名的集合。
任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
这称为Structural Typing。
所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
当然，该类型还可以有其他方法。

接口只有方法声明，没有实现，没有数据字段。
接口可以匿名嵌入其他接口，或嵌入到结构中。
对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
只有当接口存储的类型和对象都为nil时，接口才等于nil。
接口调用不会做receiver的自动转换。
接口同样支持匿名字段方法。
接口也可实现类似OOP中的多态。
空接口可以作为任何类型数据的容器。
一个类型可实现多个接口。
接口命名习惯以 er 结尾。

每个接口由数个方法组成，接口的定义格式如下：
 type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }
其中：
1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
*/

import (
	"fmt"
)

type Barker interface {
	bark()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) bark() {
	fmt.Println("汪汪汪")
}

func (cat Cat) bark() {
	fmt.Println("喵喵喵")
}

// 实现类似于JAVA多态的场景
func testBaseInterface() {
	var barker Barker

	cat := Cat{}

	dog := Dog{}

	barker = cat
	barker.bark()
	barker = dog
	barker.bark()
}

/*
喵喵喵
汪汪汪
*/

// 值接收者和指针接收者实现接口的区别,都是可以的,go的语法糖
type Mover interface {
	move()
}

func (Dog) move() {
	fmt.Println("狗会爬")
}

func (*Cat) move() {
	fmt.Println("猫会跑")
}

func TestInterfaceReceiverType() {
	var x Mover
	dog := Dog{}
	cat := Cat{}
	x = dog
	x.move()

	x = &cat
	x.move()
}

/*
狗会爬
猫会跑
*/

// 一个类型实现的接口是相互独立的，一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type Dryer struct{}

// 实现WashingMachine接口的dry()方法
func (dryer Dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type Haier struct {
	Dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (haier Haier) wash() {
	fmt.Println("洗刷刷")
}

// 接口与接口间可以通过嵌套创造出新的接口。
// 接口类型里面可以有使用其他接口进行组合,实现里面的方法即可
type Animal interface {
	Barker
	Mover
}

type Pig struct {
}

func (Pig) bark() {
	fmt.Println("哼哼哼")
}

func (Pig) move() {
	fmt.Println("猪会跑")
}

// testComposeInterface 是用来测试组合接口的方法,说明接口是可以组合的
func TestComposeInterface() {
	var animal Animal
	pig := Pig{}
	animal = pig
	animal.move()
	animal.bark()
}

/*
猪会跑
哼哼哼
*/

// 空接口的应用,十分类似于object
//1、空接口作为函数的参数,可以接收任意类型的函数参数
//2、空接口作为map的值，实现可以保存任意值的字典
func TestEmptyInterface() {
	student := make(map[string]interface{})
	student["name"] = "李白"
	student["age"] = 18
	student["married"] = false
	fmt.Println(student)
}

/*
map[age:18 married:false name:李白]
*/

// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
//想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
//              x.(T)
//其中：
// x：表示类型为interface{}的变量
// T：表示断言x可能是的类型。
func justifyType(d interface{}) {
	switch d.(type) {
	case Animal:
		fmt.Println("d is a Animal")
	case Mover:
		fmt.Println("d can only move")
	case Barker:
		fmt.Println("d can only bark")
	default:
		fmt.Println("unsupport type！")
	}
}

func TestType() {
	var pig interface{}
	pig = Pig{}
	v, ok := pig.(Animal)
	if ok {
		v.bark()
	} else {
		fmt.Println("类型断言失败")
	}

	justifyType(pig)
}
/*
哼哼哼
d is a Animal
 */