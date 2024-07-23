package main

import (
    "fmt"
)

type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof! My name is " + d.Name
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow! My name is " + c.Name
}

func makeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

// func main() {
//     dog := Dog{Name: "Buddy"}
//     cat := Cat{Name: "Whiskers"}

//     makeItSpeak(dog)
//     makeItSpeak(cat)
// }