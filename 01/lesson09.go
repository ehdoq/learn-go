package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [5]int
	a[0] = 5
	a[1] = -6
	a[2] = 11
	a[3] = 20
	a[4] = -2

	fmt.Println("Length : ", len(a))

	fmt.Println(a)

	fmt.Println("List of Elements")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d the element of a is %d\n", i, a[i])
	}

	fmt.Println("\nList of Elements")
	for index, value := range a {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}

	//-------------------------------------------------------------------

	var a2 = [5]int{2, 4, 6, 8, 10}

	fmt.Println("Length : ", len(a2))

	fmt.Println(a2)

	fmt.Println("List of Elements")
	for i := 0; i < len(a2); i++ {
		fmt.Printf("%d the element of a is %d\n", i, a[i])
	}

	fmt.Println("\nList of Elements")
	for index, value := range a2 {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}

	//-------------------------------------------------------------------

	a3 := [5]int{2, 4, 6, 8, 10}

	fmt.Println("Length : ", len(a3))

	fmt.Println(a3)

	fmt.Println("List of Elements")
	for i := 0; i < len(a3); i++ {
		fmt.Printf("%d the element of a is %d\n", i, a3[i])
	}

	fmt.Println("\nList of Elements")
	for index, value := range a {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}

	//-------------------------------------------------------------------

	a4 := [...]int{2, 4, 6, 8, 10}

	fmt.Println("Length : ", len(a4))

	fmt.Println(a4)

	fmt.Println("List of Elements")
	for i := 0; i < len(a4); i++ {
		fmt.Printf("%d the element of a is %d\n", i, a4[i])
	}

	fmt.Println("\nList of Elements")
	for index, value := range a4 {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}

	//-------------------------------------------------------------------

	a5 := [7]int{2, 4, -6, 8, -10, 24, 19}

	fmt.Println("List of Elements")
	printArray1(a5)

	fmt.Println("List of Elements")
	printArray2(a5)

	//-------------------------------------------------------------------

	var a6 = [5]int{11, -4, 7, 8, -10}
	result1, result2, result3, result4, result5 := SumArray(a6)
	fmt.Println("Sum of elements: ", result1)
	fmt.Println("Sum of positive  elements: ", result2)
	fmt.Println("Sum of negative elements : ", result3)
	fmt.Println("Sum of even elements: ", result4)
	fmt.Println("Sum of odd elements: ", result5)

	//-------------------------------------------------------------------

	var a7 = [5]int{11, -4, 7, 8, -10}
	min, max := findMinAndMax(a7)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

	//-------------------------------------------------------------------

	var names = [5]string{"mary", "peter", "kevin", "anna", "johny"}

	result6 := SortStringASC(names)
	fmt.Println("Sort ASC")
	fmt.Println(result6)

	result7 := SortStringDESC(names)
	fmt.Println("Sort DESC")
	fmt.Println(result7)

	//-------------------------------------------------------------------

	var a8 = [5]int{11, -4, 7, 8, -10}

	result8 := SortIntASC(a8)
	fmt.Println("Sort ASC")
	fmt.Println(result8)

	result9 := SortIntDESC(a8)
	fmt.Println("Sort DESC")
	fmt.Println(result9)

	//-------------------------------------------------------------------

	var names2 = [5]string{"mary", "peter", "kevin", "anna", "johny"}

	result10 := SortStringASCWithCompare(names2)
	fmt.Println("Sort ASC")
	fmt.Println(result10)

	result11 := SortStringDESCWithCompare(names2)
	fmt.Println("Sort DESC")
	fmt.Println(result11)
}

func printArray1(a [7]int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d the element of a is %d\n", i, a[i])
	}
}

func printArray2(a [7]int) {
	for index, value := range a {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}
}

func SumArray(a [5]int) (int, int, int, int, int) {
	result1 := 0
	result2 := 0
	result3 := 0
	result4 := 0
	result5 := 0
	for _, value := range a {
		result1 += value
		if value > 0 {
			result2 += value
		}
		if value < 0 {
			result3 += value
		}
		if value%2 == 0 {
			result4 += value
		}
		if value%2 != 0 {
			result5 += value
		}
	}
	return result1, result2, result3, result4, result5
}

func findMinAndMax(a [5]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func SortStringASC(a [5]string) [5]string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if strings.Compare(a[i], a[j]) > 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func SortStringDESC(a [5]string) [5]string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if strings.Compare(a[i], a[j]) < 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func SortIntASC(a [5]int) [5]int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] >= a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func SortIntDESC(a [5]int) [5]int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] <= a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func SortStringASCWithCompare(a [5]string) [5]string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if strings.Compare(a[i], a[j]) > 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func SortStringDESCWithCompare(a [5]string) [5]string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if strings.Compare(a[i], a[j]) < 0 {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
