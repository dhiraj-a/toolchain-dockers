package main

import (
	"fmt"
	"math"
	"math/rand"
)

func adderSubtractor(first, second int) (int, int) {
	return first + second, first - second
}

func returnDate() (year, month, day int, dayOfWeek string) {
	year = 2018
	month = 3
	day = 31
	dayOfWeek = "Madeupday"
	return
}

func main() {

	aString := fmt.Sprintln("My favorite number is ", rand.Intn(10), "or it could be", math.Pi)
	fmt.Println(aString)
	sum, diff := adderSubtractor(31, 45)
	fmt.Println(sum, diff)
	//var an_int, a_string, a_bool = 1, "str", true
	var year, month, day int
	var dayOfWeek string
	year, month, day, dayOfWeek = returnDate()
	fmt.Println(dayOfWeek, ",", year, "-", month, "-", day)

}
