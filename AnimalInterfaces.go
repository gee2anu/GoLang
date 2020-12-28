package main

import (
	"fmt"
	"strings"
)

/* Write a program which allows the user to create a set of animals and to get
information about those animals. Each animal has a name and can be either a cow,
bird, or snake. With each command, the user can either create a new animal of one
of the three types, or the user can request information about an animal that he/she
has already created. Each animal has a unique name, defined by the user. Note that
the user can define animals of a chosen type, but the types of animals are restricted
to either cow, bird, or snake. The following table contains the three types of animals
and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.*/

type Cow struct{ name string }

func (c Cow) Eat()   { fmt.Println("Grass") }
func (c Cow) Move()  { fmt.Println("Walk") }
func (c Cow) Speak() { fmt.Println("mooo") }

type Bird struct{ name string }

func (b Bird) Eat()   { fmt.Println("Worms") }
func (b Bird) Move()  { fmt.Println("Fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{ name string }

func (s Snake) Eat()   { fmt.Println("Mice") }
func (s Snake) Move()  { fmt.Println("Slither") }
func (s Snake) Speak() { fmt.Println("hssss") }

type Animal interface {
	Eat()
	Move()
	Speak()
}

var animalMap = make(map[string]Animal)

func addAnimal(species string, name string) {
	switch species {
	case "cow":
		c1 := Cow{name}
		animalMap[name] = c1
	case "bird":
		b1 := Bird{name}
		animalMap[name] = b1
	case "snake":
		s1 := Snake{name}
		animalMap[name] = s1
	}
	return
}

func outputAnimal(name string, prop string) {
	switch prop {
	case "eat":
		animalMap[name].Eat()
	case "speak":
		animalMap[name].Speak()
	case "move":
		animalMap[name].Move()
	}
	return
}

func main() {
	var inp string
	var name string
	var prop string
	for {
		fmt.Printf("\n Enter input:(Note : Invalid input exits the program!) ")
		fmt.Scanf("%s %s %s", &inp, &name, &prop)

		if inp == "newanimal" {
			addAnimal(strings.ToLower(prop), name)
			inp = ""
			fmt.Println("Created it!!")
		} else if inp == "query" {
			outputAnimal(name, strings.ToLower(prop))
		} else {
			break
		}
	}
	return
}
