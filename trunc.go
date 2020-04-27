package main
import "fmt"
import "math"

//Week 2 Assignment 1 .
//Question : Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

func main() {
var f_input float64
fmt.Printf("Enter a floating point number : ")
fmt.Scan(&f_input)
fmt.Printf("%d \n", math.Trunc(f_input))
}
