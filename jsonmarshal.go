package main

import (
  "fmt"
  "bufio"
  "os"
  "encoding/json"
)

//Week 4  : Assignment 1 : Question
//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
//Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

//It would be good if you read up on Bufio at this point.

func main(){
  var name, addr string
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Printf("Enter a name : \t ")
  scanner.Scan()
  name = scanner.Text()
  fmt.Printf("Enter a address : \t ")
  scanner.Scan()
  addr = scanner.Text()

  tmap  := map[string]string{"name":name, "address":addr}

  barr, err := json.Marshal(tmap)
  if err == nil {
    fmt.Println(string(barr))
  }
}
