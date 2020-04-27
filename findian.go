package main

import "fmt"
import "bufio"
import "os"
import "strings"

//Week 2 Assignment 2 : Question
//Write a program which prompts the user to enter a string.
//The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
//The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise.
//The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
//Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
//The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.


//The tricky part here was to get the input with spaces and process it. It became easy once I found out about the bufio package.

func main() {

  fmt.Printf("\n Enter a string : \t")
  inp, err := bufio.NewReader(os.Stdin).ReadString('\n')

  line := strings.ToLower(inp)

  if err == nil {
    fmt.Printf("\n string is : %s and length is %d \n", line, len(line))
  }
  i_pos := 0
  n_pos := len(line) - 2
  a_count := 0
  fmt.Printf("i_pos %d, n_pos %d \n", i_pos, n_pos)
  for i:= 0; i < len(line) - 1; i++ {
    pos := i
    ch := line[i]
    fmt.Printf("\n Char at %d position is %c \n", pos, ch)
    if i_pos == pos {
      if ch != 'i'{
        fmt.Printf("Not Found!\n")
        return
      }
    }else if n_pos == pos{
      if ch != 'n'{
        fmt.Printf("Not Found!\n")
        return
      }
    }
    if a_count <= 0 {
      if(ch == 'a'){
        a_count = a_count + 1
      }
    }

  }
  fmt.Printf("a_count is %d \n", a_count)
  if a_count > 0 {
    fmt.Printf("Found!\n")
  }else {
    fmt.Printf("Not Found!\n")
  }
}
