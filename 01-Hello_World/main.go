// Package main is a special Go package that tells the compiler this will contain the main function, as seen below.
package main

// Import is a key word that is used to add packages to your program.  Think of them like moving boxes that you have written the contents on the box.  If you want to use the screwdriver you first need to look for the box labeled tools and then find the screwdriver
import (
	// This imports the "box" with the fmt tools that allows you to print characters to the screen.
	"fmt"
)

// This is the main method/function.  When your program runs this will be the starting point.
func main() {
	// First we tell Go we want to open the "fmt" box and the the "Println" tool out.  Println will print any string that happens to be in it.
	fmt.Println("Hello World!")
}