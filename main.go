package main

import (
	"fmt"
)

func greeting(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	fmt.Println("This is awesome project build on GoLang")
	name := "eddie"
	greeting(name)
}
