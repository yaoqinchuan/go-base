package base

import "fmt"

func testBasic() {
	arrayA := [2]int{100, 200}
	testArray(&arrayA)
	arrayB := arrayA[:]
	testSlice(&arrayB)

	fmt.Printf("func array : %p, %v\n", &arrayA, arrayA)
}

func testArray(x *[2]int) {
	fmt.Printf("func array : %p, %v\n", x, *x)
	(*x)[1] += 100
}
func testSlice(x *[]int) {
	fmt.Printf("func array : %p, %v\n", x, *x)
	(*x)[1] += 100
}

