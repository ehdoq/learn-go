package main

import "fmt"

func main() {
	PointerDemo1()
	PointerDemo2()
	PointerDemo3()

	//-------------------------------------------------------------------

	var a int = 1
	ChangeValue1(a)
	fmt.Println("a:", a)
	ChangeValue2(&a)
	fmt.Println("a:", a)

	b, c := 3, 4
	p, q := &b, &c
	Swap1(b, c)
	fmt.Println("b:", b, ", c:", c)
	Swap2(p, q)
	fmt.Println("b:", b, ", c:", c)

	//-------------------------------------------------------------------

	var a2 = []int{5, 9, 1, 2, 8}
	p2 := &a2

	fmt.Println("Element 0: ", (*p2)[0])

	fmt.Println("List of Elements")
	for i := 0; i < len(a2); i++ {
		fmt.Print((*p2)[i], "  ")
	}

	fmt.Println("\nList of Elements")
	for index, value := range *p2 {
		fmt.Println(index, ": ", value)
	}

	fmt.Println("Elements of A array")
	ModifyArray(p2)
	fmt.Println(a2)

}

func PointerDemo1() {
	var a int = 5
	var p *int = &a
	fmt.Println("Value of a variable: ", a)
	fmt.Println("Address of a variable: ", &a)
	fmt.Println("Address of p pointer: ", p)
	fmt.Println("Value of p pointer: ", *p)

	*p = 6
	fmt.Println("Value of a variable: ", a)
	fmt.Println("Address of a variable: ", &a)
	fmt.Println("Address of p pointer: ", p)
	fmt.Println("Value of p pointer: ", *p)
}

func PointerDemo2() {
	var a int = 5
	p := &a
	fmt.Println("Value of a variable: ", a)
	fmt.Println("Address of a variable: ", &a)
	fmt.Println("Address of p pointer: ", p)
	fmt.Println("Value of p pointer: ", *p)

	*p = 6
	fmt.Println("Value of a variable: ", a)
	fmt.Println("Address of a variable: ", &a)
	fmt.Println("Address of p pointer: ", p)
	fmt.Println("Value of p pointer: ", *p)
}

func PointerDemo3() {
	var a int = 5
	var b int = 10
	p, q := &a, &b

	fmt.Println("Value of a variable: ", a)
	fmt.Println("Address of a variable: ", &a)
	fmt.Println("Value of b variable: ", b)
	fmt.Println("Address of b variable: ", &b)

	fmt.Println("Address of p pointer: ", p)
	fmt.Println("Value of p pointer: ", *p)
	fmt.Println("Address of q pointer: ", q)
	fmt.Println("Value of q pointer: ", *q)
}

func ChangeValue1(a int) {
	a = 2
}

func ChangeValue2(p *int) {
	*p = 2
}

func Swap1(a, b int) {
	temp := a
	a = b
	b = temp
}

func Swap2(p, q *int) {
	temp := *p
	*p = *q
	*q = temp
}

func ModifyArray(p *[]int) {
	(*p)[2] = 22
}
