package main

import "fmt"

type Person struct {
    firstName string
    lastName string
    age int
}

func main() {
    //We can initialize a struct using various methods
    
    //1. Using var keyword
    var p1 Person 
    p1.firstName = "sneha"
    p1.lastName = "kumari"
    p1.age = 25
    
    fmt.Println("FirstName:",p1.firstName)
    fmt.Println("LastName:",p1.lastName)
    fmt.Println("Age:",p1.age)
    
    //2.Using a struct literal
    p2 := Person {
        firstName : "Sanjay",
        lastName : "Kumar",
        age : 30,
    }
    fmt.Println("FirstName:",p2.firstName)
    fmt.Println("LastName:",p2.lastName)
    fmt.Println("Age:",p2.age)
    
    //3.Using new keyword
    p3 := new(Person)
    p3.firstName = "Ishaan"
    p3.lastName = "choudhary"
    p3.age = 1
    
    fmt.Println("FirstName:",p3.firstName)
    fmt.Println("LastName:",p3.lastName)
    fmt.Println("Age:",p3.age)
}
