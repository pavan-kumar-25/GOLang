package main

import (
	"fmt"
)

// Boat represents a simple structure for a boat
type Boat struct {
	name   string
	capacity int
}

// navigateByValue demonstrates call by value
func navigateByValue(boat Boat) {
	boat.capacity += 10 // Modifying the capacity locally
	fmt.Printf("Inside navigateByValue: Boat name: %s, Capacity: %d\n", boat.name, boat.capacity)
}

// navigateByPointer demonstrates call by pointer
func navigateByPointer(boat *Boat) {
	boat.capacity += 10 // Modifying the capacity through a pointer
	fmt.Printf("Inside navigateByPointer: Boat name: %s, Capacity: %d\n", boat.name, boat.capacity)
}

// navigateByFunction demonstrates call by function
func navigateByFunction(boatFunc func(Boat)) {
	boat := Boat{name: "Ferry", capacity: 100}
	boatFunc(boat) // Passing a function as an argument
	fmt.Printf("Inside navigateByFunction: Boat name: %s, Capacity: %d\n", boat.name, boat.capacity)
}

func main() {
	boat := Boat{name: "Rowboat", capacity: 50}

	// Call by value
	fmt.Println("Before navigateByValue: Boat name:", boat.name, "Capacity:", boat.capacity)
	navigateByValue(boat)
	fmt.Println("After navigateByValue: Boat name:", boat.name, "Capacity:", boat.capacity)

	// Call by pointer
	fmt.Println("Before navigateByPointer: Boat name:", boat.name, "Capacity:", boat.capacity)
	navigateByPointer(&boat)
	fmt.Println("After navigateByPointer: Boat name:", boat.name, "Capacity:", boat.capacity)

	// Call by function
	fmt.Println("Before navigateByFunction: Boat name:", boat.name, "Capacity:", boat.capacity)
	navigateByFunction(navigateByValue)
	fmt.Println("After navigateByFunction: Boat name:", boat.name, "Capacity:", boat.capacity)
}

// output:
// Before navigateByValue: Boat name: Rowboat Capacity: 50
// Inside navigateByValue: Boat name: Rowboat, Capacity: 60
// After navigateByValue: Boat name: Rowboat Capacity: 50
// Before navigateByPointer: Boat name: Rowboat Capacity: 50
// Inside navigateByPointer: Boat name: Rowboat, Capacity: 60
// After navigateByPointer: Boat name: Rowboat Capacity: 60
// Before navigateByFunction: Boat name: Rowboat Capacity: 60
// Inside navigateByValue: Boat name: Ferry, Capacity: 110
// Inside navigateByFunction: Boat name: Ferry, Capacity: 100
// After navigateByFunction: Boat name: Rowboat Capacity: 60