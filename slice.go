package main

import "fmt"
import "sort"
import "strconv"

//Week 3 : Assignment 1 : Question
// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
//The program should be written as a loop.
//Before entering the loop, the program should create an empty integer slice of size (length) 3.
//During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
//The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
//The slice must grow in size to accommodate any number of integers which the user decides to enter.
//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

//I used in built sort function in this program. As the list is always maintained in sorted order, inserting at the relevant position would keep the list sorted always.
//This could be taken as an optimisation.

func main() {
  var inp string
  slice := make([]int, 0, 3)
  fmt.Printf("\n Enter 'X' to exit \n")
  for{
    fmt.Printf("\n Enter an integer: ")
    fmt.Scanf("%s", &inp)
    if inp == "X"{
      break
    }
    i, err := strconv.Atoi(inp)
    if err == nil {
    slice = append(slice, i)
   }

    sort.Ints(slice)
    fmt.Printf("\n The sorted slice is : \n")
    fmt.Println(slice)
  }

}
