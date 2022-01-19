package main

import "fmt"

func main() {
	var name1 string = "Yonatan" //this is a full variable decleration plus initlization. it specifies var(keyword),variable name and type.
	var name2 string             //variable decleration without initlization

	name2 = "Maor"
	name3 := "Avi" //short variable decleration

	name3, name4 := "Avi wizman", "Yosef Chaim" //short variable assignment to two variables at once. atleast one of them has to be not perviously declared

	p := &name3     //p is a pointer that stores name3 adress
	fmt.Println(*p) //this will print out the value pointed by p which is avi wizman

}
