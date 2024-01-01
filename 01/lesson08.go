package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcDEF123"
	fmt.Println("Length: ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%5c", s[i])
	}

	//-------------------------------------------------------------------

	s1 := "abc"
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)

	s2 := "DEF"
	s2 = strings.ToLower(s2)
	fmt.Println(s2)

	//-------------------------------------------------------------------

	s3 := "computer"
	s4 := "pu"

	result1 := strings.HasPrefix(s3, s4)
	fmt.Println("result 1: ", result1)

	result2 := strings.HasSuffix(s3, s4)
	fmt.Println("result 2: ", result2)

	result3 := strings.Contains(s3, s4)
	fmt.Println("result 3: ", result3)

	//-------------------------------------------------------------------

	name1 := "mary"
	name2 := "peter"
	name3 := "anna"
	name4 := "mary"

	result4 := strings.Compare(name1, name2)
	fmt.Println("result 4:", result4)

	result5 := strings.Compare(name1, name3)
	fmt.Println("result 5:", result5)

	result6 := strings.Compare(name1, name4)
	fmt.Println("result 6:", result6)

	//-------------------------------------------------------------------

	s5 := "p01,name 1,4.5,20"
	result := strings.Split(s5, ",")
	for _, value := range result {
		fmt.Println(value)
	}
}
