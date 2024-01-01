package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a = [5]int{4, 1, -2, 9, 10}

	fmt.Println("a: ", a)

	slice1 := a[0:2]
	fmt.Println("slice 1: ", slice1)

	slice1[1] = 999
	fmt.Println("a: ", a)
	fmt.Println("slice 1: ", slice1)

	slice2 := a[:2]
	fmt.Println("slice 2: ", slice2)

	slice3 := a[2:]
	fmt.Println("slice 3: ", slice3)

	slice4 := a[:]
	fmt.Println("slice 4: ", slice4)

	slice5 := []int{7, 2, 4, 9}
	fmt.Println("slice 5: ", slice5)

	slice6 := []string{"name 1", "name 2", "name 3", "name 4"}
	fmt.Println("slice 6: ", slice6)

	//-------------------------------------------------------------------

	var a2 = [5]int{4, 1, -2, 9, 10}

	fmt.Println("a: ", a2)
	fmt.Println("Length of a2: ", len(a2))

	slice7 := a[0:2]
	fmt.Println("slice 1: ", slice7)

	slice1 = append(slice7, 25)
	fmt.Println("slice 7: ", slice7)
	fmt.Println("a2: ", a2)
	fmt.Println("Length of a2: ", len(a2))

	slice1 = append(slice7, 100, 200, 300)
	fmt.Println("slice 7: ", slice7)
	fmt.Println("a2: ", a2)
	fmt.Println("Length of a2: ", len(a2))

	//-------------------------------------------------------------------

	//You can directly append one slice to another using the "..." operator.
	//This operator expands the slice to a list of arguments.

	slice8 := []string{"Jack", "John", "Peter"}
	slice9 := []string{"Bill", "Mark", "Steve"}

	slice10 := append(slice8, slice9...)

	fmt.Println("slice8 = ", slice8)
	fmt.Println("slice9 = ", slice9)
	fmt.Println("After appending slice8 & slice9 = ", slice10)

	//-------------------------------------------------------------------

	var a3 = [5]int{4, 1, -2, 9, 10}

	slice11 := a3[0:2]
	fmt.Println("slice 11")
	fmt.Println("\tLength:", len(slice11))
	fmt.Println("\tCapacity:", cap(slice11))

	slice12 := a3[:4]
	fmt.Println("slice 12")
	fmt.Println("\tLength:", len(slice12))
	fmt.Println("\tCapacity:", cap(slice12))

	slice13 := a3[4:]
	fmt.Println("slice 13")
	fmt.Println("\tLength:", len(slice13))
	fmt.Println("\tCapacity:", cap(slice13))

	slice14 := a3[:]
	fmt.Println("slice 14")
	fmt.Println("\tLength:", len(slice14))
	fmt.Println("\tCapacity:", cap(slice14))

	//-------------------------------------------------------------------

	slice15 := make([]int, 3)
	fmt.Println("Length:", len(slice15))
	fmt.Println("Capacity:", cap(slice15))
	slice15[0] = 4
	slice15[1] = 2
	slice15[2] = -4
	slice15 = append(slice15, 20)
	slice15 = append(slice15, 20, 30, 40, 50)
	fmt.Println("slice 15: ", slice15)

	slice16 := make([]int, 4, 8)
	fmt.Println("Length:", len(slice16))
	fmt.Println("Capacity:", cap(slice16))
	slice16[0] = 2
	slice16[1] = 9
	slice16[2] = -8
	fmt.Println("slice 16: ", slice16)

	//-------------------------------------------------------------------

	a4 := [7]int{2, 4, -6, 8, -10, 24, 19}

	slice17 := a4[:]

	fmt.Println("List of Elements")
	printSlice1(slice17)

	fmt.Println("List of Elements")
	printSlice2(slice17)

	//-------------------------------------------------------------------

	var a5 = [3]int{4, 1, -2}

	ChangeValueWithArray(a5)
	fmt.Println("Change Value with Array")
	fmt.Println("a5: ", a5)

	fmt.Println("Change Value with Slice")
	ChangeValueWithSlice(a5[:])
	fmt.Println("a5: ", a5)

	//-------------------------------------------------------------------

	fmt.Println("Sort ASC")
	slice18 := []int{4, 1, -2, 9, 10}
	sort.Ints(slice18)
	fmt.Println(slice18)

	fmt.Println("Sort DESC")
	slice19 := []int{4, 1, -2, 9, 10}
	sort.Slice(slice19, func(i, j int) bool {
		return slice19[i] >= slice19[j]
	})
	fmt.Println(slice19)

	//-------------------------------------------------------------------

	fmt.Println("Sort ASC")
	names1 := []string{"mary", "peter", "kevin", "anna", "johny"}
	sort.Strings(names1)
	fmt.Println(names1)

	fmt.Println("Sort DESC")
	names2 := []string{"mary", "peter", "kevin", "anna", "johny"}
	sort.Slice(names2, func(i, j int) bool {
		return strings.Compare(names2[i], names2[j]) > 0
	})
	fmt.Println(names2)
}

func printSlice1(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d the element of a is %d\n", i, slice[i])
	}
}

func printSlice2(slice []int) {
	for index, value := range slice {
		fmt.Printf("%d the element of a is %d\n", index, value)
	}
}

func ChangeValueWithArray(a [3]int) {
	a[1] = 999
}

func ChangeValueWithSlice(slice []int) {
	slice[1] = 999
}
