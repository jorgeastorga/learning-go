package main

import "fmt"

func main() {

	var x int = 1
	fmt.Println(x)

	var p *int = &x
	fmt.Println(p)
	fmt.Printf("This is the type for p: %T\n", p)
	fmt.Printf("Value in default format for p: %v\n", p)
	fmt.Printf("Pointer notation (base 16 notation with leading 0x) for p: %p\n", p)

	fmt.Println(*p)

	increment(p)
	fmt.Printf("Value of x after increment is called: %d\n", x)

	pAlt := p
	fmt.Printf("Value in default format for pAlt: %v\n", pAlt)

	increment(&x)

}

func increment(number *int) {
	fmt.Printf("Address of number: %p\n", number)
	fmt.Printf("Value of number before incrementing: %d\n", *number)
	*number++
	fmt.Printf("Value of number after icrementing: %d\n", *number)
}
