package main

import "fmt"

func main() {
	// print a simple string, no newline
	fmt.Print("Welcome to go!")

	// print new line
	fmt.Println("This string ends with a newline")

	// print a string with values, note how go adds space between the string and value
	const answer = 42
	fmt.Println("The answer to life is", answer)

	// print a string with multiple interspersed values
	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "you get", c)

	// print a slice of data
	items := []int{10, 20, 40, 80}
	length, arr := fmt.Println(items)
	fmt.Println(length, arr)

	// *** Formatting ***
	x := 20
	f := 123.45

	// basic formatting
	fmt.Printf("%d\n", x) // Decimal representation
	fmt.Printf("%x\n", x) // Hexadecimal representation
	fmt.Printf("%b\n", x) // Binary representation
	fmt.Printf("%o\n", x) // Octal representation
	fmt.Printf("%c\n", x) // Character representation

	// boolean can be printed as true or false
	fmt.Printf("%t\n", x > 10)

	fmt.Printf("%f\n", f) // Floating point representation
	fmt.Printf("%e\n", f) // Exponential representation

	// using explicit argument  indexes
	fmt.Printf("%[2]d %[1]d\n", 52, 40)  // Print second arg first
	fmt.Printf("%d %#[1]o %#[1]x\n", 52) // Print same value in different formats, # for go syntax format

	circle1 := Circle{
		radius: 10,
		border: 1,
	}

	// print a value in default format
	fmt.Printf("%v\n", circle1)  // just values
	fmt.Printf("%+v\n", circle1) // With field names
	fmt.Printf("%T\n", circle1)  // type

	// Sprintf is same as Printf but returns a string
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)

	// *** Advanced Formatting ***
	fa := 123.4567

	// control the decimal precision
	fmt.Printf("%.2f\n", fa)

	// print with width of 10 and default precision
	fmt.Printf("%10f\n", fa)

	// print with padding and precision
	fmt.Printf("%10.2f\n", fa)

	// always use a + sign
	fmt.Printf("%+10.2f\n", fa)

	// pad with 0s instead of spaces
	fmt.Printf("%010.2f\n", fa)
}

type Circle struct {
	radius int
	border int
}
