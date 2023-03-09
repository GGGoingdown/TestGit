package main

import (
	"fmt"
)

func greeting(name string, age int) string {
	return fmt.Sprintf("Hello %s. You are %d years old", name, age)
}

func main() {
	fmt.Println("This is awesome project build on GoLang")
	name := "eddie"
	age := 10
	greeting(name, age)
	sayHello(name)
}
