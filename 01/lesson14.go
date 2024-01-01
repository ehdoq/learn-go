package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Create new folder named data.

	file, err := os.Create("sample/a.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
		fmt.Println("Done")
	}

	//-------------------------------------------------------------------

	//Create new folder named data.
	//In data folder, create new text file named a.txt.

	file2, err2 := os.Stat("sample/a.txt")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Name: ", file2.Name())
		fmt.Println("Size(bytes): ", file2.Size())
		fmt.Println("Permission: ", file2.Mode())
		fmt.Printf("Perm: %04o", file2.Mode().Perm())
	}

	//-------------------------------------------------------------------
	//Create new folder named data.
	//In data folder, create new text file named a.txt.

	err3 := os.Remove("data/a.txt")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("Done")
	}

	//-------------------------------------------------------------------

	//Create new folder named data.
	//In data folder, create new text file named a.txt.
	//Write Hello World in a.txt.

	result, err4 := os.ReadFile("sample/a.txt")

	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(string(result))
	}

	//-------------------------------------------------------------------

	//Create new folder named data.
	//In data folder, create new text file named a.txt.
	/*
		Line 1
		Line 2
		Line 3
		Line 4
	*/

	file3, err5 := os.Open("sample/a.txt")
	if err5 != nil {
		fmt.Println(err5)
	} else {
		scanner := bufio.NewScanner(file3)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}
	file3.Close()

	//-------------------------------------------------------------------

	//Create new folder named data.
	//In data folder, create new text file named product.csv.
	/*
		p01,name 1,4.5,20
		p02,name 2,7,11
		p03,name 3,2,8
	*/

	file4, err6 := os.Open("sample/product.csv")
	if err6 != nil {
		fmt.Println(err6)
	} else {
		scanner := bufio.NewScanner(file4)
		for scanner.Scan() {
			line := scanner.Text()
			result := strings.Split(line, ",")
			fmt.Println("id: ", result[0])
			fmt.Println("name: ", result[1])
			fmt.Println("price: ", result[2])
			fmt.Println("quantity: ", result[3])
			fmt.Println("-------------------")
		}
	}
	file4.Close()

	//-------------------------------------------------------------------

	//Create new folder named data.

	file5, err7 := os.Create("sample/b.txt")
	if err7 != nil {
		fmt.Println(err7)
	} else {
		file5.WriteString("Hello World")
		fmt.Println("Done")
	}
	file5.Close()

	//-------------------------------------------------------------------

	//Create new folder named data.
	//In data folder, create new text file named a.txt.
	//Hello World

	file6, err8 := os.OpenFile("sample/a.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err8 != nil {
		fmt.Println(err8)
	} else {
		fmt.Fprintln(file6, "Hi")
		fmt.Println("Done")
	}
	file6.Close()
}
