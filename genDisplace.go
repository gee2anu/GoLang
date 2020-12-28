package main

import (
   "fmt"
   "math"
 )

//Week 2 : Assignment 1 : // QUESTION:
/* Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =(1/2)a * pow(t,2) + v * t + s

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))*/

func GenDisplaceFn(acc, v0, s0 float64) func(float64) float64 {
  fn := func (t float64) float64 {
    return ((0.5 * acc * math.Pow(t , 2)) + (v0 * t) + (s0))
  }
  return fn
}

func main(){
  var acceleration , initial_velocity, initial_displacement float64
  var time float64

  fmt.Printf("\n Enter the acceleration : ")
  fmt.Scanf("%f", &acceleration)

  fmt.Printf("\n Enter the Initial Velocity : ")
  fmt.Scanf("%f", &initial_velocity)

  fmt.Printf("\n Enter the Initial Displacement : ")
  fmt.Scanf("%f", &initial_displacement)

  fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)

  fmt.Printf("\n Enter the time : ")
  fmt.Scanf("%f", &time)

  fmt.Println("\n Displacement after ", time, " seconds is ", fn(time))


}
