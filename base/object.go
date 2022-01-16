package base

import "fmt"

type Person struct {
	name string
	sex string
	age int32
}

type Student struct {
	*Person
	id int32
	name string
}

// 访问匿名类的内部字段可以使用类
func TestObjectAnonymousField()  {
	var student Student
	student = Student{
		&Person{"张三", "男", 32}, 10012, "张三",
	}
	student.name = "李四"
	student.Person.name = "王五"
	fmt.Println(student.Person)
	fmt.Println(student)
}
/*
&{王五 男 32}
{0xc000062330 10012 李四}
 */