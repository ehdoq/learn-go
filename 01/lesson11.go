package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println("Today:", today)
	fmt.Println("Year:", today.Year())

	month := today.Month()
	fmt.Println("Month: ", month)
	fmt.Println("Month: ", int(month))
	fmt.Println("Day: ", today.Day())
	fmt.Println("Hour: ", today.Hour())
	fmt.Println("Minutes: ", today.Minute())
	fmt.Println("Second: ", today.Second())

	//-------------------------------------------------------------------

	today2 := time.Now()
	fmt.Println("Today: ", today2.Format("02/01/2006 15:04:05"))

	str := "27/12/2018"
	date, error := time.Parse("02/01/2006", str)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(date.Format("2006-01-02"))
	}

	//-------------------------------------------------------------------

	DateDemo1()
	DateDemo2()
	DateDemo3()
}

func DateDemo1() {
	today := time.Now()

	fmt.Println("Add Days")
	time1 := today.AddDate(0, 0, 2)
	fmt.Println("time 1:", time1.Format("02/01/2006"))

	fmt.Println("Add Months")
	time2 := today.AddDate(0, 3, 0)
	fmt.Println("time 2:", time2.Format("02/01/2006"))

	fmt.Println("Add Years")
	time3 := today.AddDate(4, 0, 0)
	fmt.Println("time 3:", time3.Format("02/01/2006"))

	time4 := today.Add(2 * 24 * time.Hour)
	fmt.Println("time 4:", time4.Format("02/01/2006"))
}

func DateDemo2() {
	strDate1 := "20/10/2018"
	date1, _ := time.Parse("02/01/2006", strDate1)
	strDate2 := "25/10/2018"
	date2, _ := time.Parse("02/01/2006", strDate2)
	result1 := date2.Sub(date1).Hours()
	fmt.Println("Hours: ", result1)
	result2 := date2.Sub(date1).Minutes()
	fmt.Println("Minutes: ", result2)
}

func DateDemo3() {
	date1 := time.Date(2018, 11, 20, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2018, 11, 25, 0, 0, 0, 0, time.UTC)

	result1 := date1.Before(date2)
	fmt.Println("Before: ", result1)

	result2 := date1.After(date2)
	fmt.Println("After: ", result2)

	result3 := date1.Equal(date2)
	fmt.Println("Equal: ", result3)
}
