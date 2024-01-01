package main

import (
	"fmt"
)

func main() {
	Hello()
	Sum(1, 2)
	Sub(10, 5)

	//-------------------------------------------------------------------

	result1 := Hi("MG")
	fmt.Println("result 1: ", result1)

	result2 := Sum1(1, 2)
	fmt.Println("result 2: ", result2)

	result3 := Sum2(1, 2)
	fmt.Println("result 3: ", result3)

	r1, r2, r3, r4 := Calculate1(10, 5)
	fmt.Println("r1:", r1, ", r2:", r2, ", r3:", r3, ", r4:", r4)

	r5, r6, r7, r8 := Calculate2(10, 5)
	fmt.Println("r5:", r5, ", r6:", r6, ", r7:", r7, ", r8:", r8)

	r9, _, r10, _ := Calculate3(10, 5)
	fmt.Println("r9:", r9, ", r10:", r10)
}

func Hello() {
	fmt.Println("Hello World")
}

func Sum(a int, b int) {
	fmt.Println("Sum: ", (a + b))
}

func Sub(a, b int) {
	fmt.Println("Sub: ", (a - b))
}

func Hi(fullName string) string {
	return fmt.Sprintf("Hi %s", fullName)
}

func Sum1(a int, b int) int {
	return a + b
}

func Sum2(a, b int) (result int) {
	result = a + b
	return result
}

func Calculate1(a, b int) (int, int, int, float32) {
	result1 := a + b
	result2 := a - b
	result3 := a * b
	result4 := float32(a) / float32(b)
	return result1, result2, result3, result4
}

func Calculate2(a int, b int) (result1 int, result2 int, result3 int, result4 float32) {
	result1 = a + b
	result2 = a - b
	result3 = a * b
	result4 = float32(a) / float32(b)
	return result1, result2, result3, result4
}

func Calculate3(a, b int) (result1, result2, result3 int, result4 float32) {
	result1 = a + b
	result2 = a - b
	result3 = a * b
	result4 = float32(a) / float32(b)
	return result1, result2, result3, result4
}
