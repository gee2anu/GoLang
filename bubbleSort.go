package main

import "fmt"

//Week 1 : Assignment 1 
///Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

///As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

///A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

func Swap(sli []int, j int){
  sli[j], sli[j+1] = sli[j+1], sli[j]
  return
}

func BubbleSort(sli []int){
  swaps := false

  for i , _ := range(sli){
    for j := 0; j < len(sli) - i - 1; j++{
      if sli[j] > sli[j + 1]{
        Swap(sli, j)
        swaps = true
      }

    }
    if swaps != true{
      break
    }
    swaps = false
  }

  return
}

func main(){

  //Declare variables
  sli := make([]int, 0, 10)
  n := 0
  temp := 0
  //Input number of integers
  fmt.Print("Enter number of integers to be sorted (upto 10)  : ")
  fmt.Scanf("%d", &n)
  if n > 10 {
    fmt.Print("%d is greater than 10. Assignment specifies only upto 10 numbers ", n)
    return
  } else if (n < 0) {
    fmt.Print("%d is lesser than 0. Invalid value!! ", n)
    return
  }
  //Input the integers
  for i:=0; i < n; i++{
    fmt.Print("Enter the number : ")
    fmt.Scanf("%d", &temp)
    sli = append(sli, temp)
  }
  fmt.Println("Before sorting : ", sli)
  BubbleSort(sli)
  fmt.Println("After sorting : ", sli)
}
