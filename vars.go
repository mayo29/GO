package main

import "fmt"

func main() {
	var name1 string = "Yonatan" //this is a full variable decleration plus initlization. it specifies var(keyword),variable name and type.
	var name2 string             //variable decleration without initlization
	fmt.Println(name1)
	fmt.Println(name2)

	name2 = "Maor"
	name3 := "Avi" //short variable decleration

	name3, name4 := "Avi wizman", "Yosef Chaim" //short variable assignment to two variables at once. atleast one of them has to be not perviously declared
	fmt.Println(name4)
	p := &name3     //p is a pointer that stores name3 adress
	fmt.Println(*p) //this will print out the value pointed by p which is avi wizman

	var pointer *int     //the zero value for a pointer is nil
	fmt.Println(pointer) //this will print out nil

	intPointer := new(int)   //creates a default int pointer
	fmt.Println(*intPointer) //will print out zero
	*intPointer = 3
	fmt.Println(*intPointer) // will print out 3

	stringPointer := new(string) //creates default string pointer
	fmt.Println(*stringPointer)  //should print nothing (empty string)
	*stringPointer = "jhon"
	fmt.Println(*stringPointer) //should print jhon

}
