/*
Scenario 1:
You've been assigned to develop a flight management system for an airline company
that operates  international flights. Your task is to create a Go program
that manages flight scheduling, passenger bookings, and fare calculations.
Your program should incorporate variables, data types, arrays, slices, maps, structures, control flow, loops, and functions.

Here's what you need to do:

Define a structure named Flight to represent each flight in the airline's schedule.
Include attributes such as FlightNumber, Departure, Destination, DepartureTime, and Capacity.

Create an array named passengers to store the details of passengers
who have booked tickets for a particular flight.

Define a slice named destinations to represent the destinations served by the airline,
including both international locations.

Implement a map named farePrices to store the fare prices for different flight routes.

Develop a function called calculateTotalFare that takes the number of passengers
and the flight route as input parameters and returns the total fare for the journey.

Use control flow statements to handle the following scenarios:

If the number of passengers exceeds the flight capacity, print a message
indicating that the flight is fully booked.
Display an error message if a passenger tries to book a ticket for
a destination that is not served by the airline.

Utilize loops to perform the following tasks:

Display the list of available flights along with their departure times and destinations.
Allow passengers to book tickets for a selected flight and add their details to the passengers array.
Calculate the total fare for a journey based on the number of passengers and the selected flight route.*/


package main

import "fmt"

// Structure to represent each flight
type Flight struct {
    FlightNumber   string
    Departure      string
    Destination    string
    DepartureTime  string
    Capacity       int
}



func main() {
    // Array to store details of passengers
    var passengers []string

    // Slice to represent destinations served by the airline
    destinations := []string{"New York", "London", "Tokyo", "Paris", "Sydney"}

    // Map to store fare prices for different flight routes
    farePrices := map[string]float64{
        "New York": 500.0,
        "London":   600.0,
        "Tokyo":    800.0,
        "Paris":    550.0,
        "Sydney":   1000.0,
    }

    // Define flights
    flights := []Flight{
        {"FL001", "New York", "London", "08:00", 200},
        {"FL002", "London", "New York", "09:00", 220},
        {"FL003", "Tokyo", "Paris", "10:00", 180},
    }

    // Display available flights
    fmt.Println("Available Flights:")
    for _, flight := range flights {
        fmt.Printf("Flight: %s, \tDeparture: %s, \tDestination: %s, \tDeparture Time: %s, \tCapacity: %d\n",
            flight.FlightNumber, flight.Departure, flight.Destination, flight.DepartureTime, flight.Capacity)
    }

    // Allow passengers to book tickets
    var flightChoice string
    fmt.Print("\nEnter the flight number you want to book: ")
    fmt.Scanln(&flightChoice)

    var numPassengers int
    fmt.Print("Enter the number of passengers: ")
    fmt.Scanln(&numPassengers)

    var destinationChoice string
    fmt.Print("Enter the destination: ")
    fmt.Scanln(&destinationChoice)

    // Check if the destination is served by the airline
    validDestination := false
    for _, dest := range destinations {
        if dest == destinationChoice {
            validDestination = true
            break
        }
    }

    if !validDestination {
        fmt.Println("Error: Destination not served by the airline.")
        return
    }

    // Find the selected flight
    var selectedFlight Flight
    for _, flight := range flights {
        if flight.FlightNumber == flightChoice {
            selectedFlight = flight
            break
        }
    }

    // Check if the flight is fully booked
    if numPassengers > selectedFlight.Capacity {
        fmt.Println("Error: Flight is fully booked.")
        return
    }

    // Book tickets
    for i := 0; i < numPassengers; i++ {
        var passengerName string
        fmt.Printf("Enter name of passenger %d: ", i+1)
        fmt.Scanln(&passengerName)
        passengers = append(passengers, passengerName)
    }

    // Calculate total fare
    totalFare := calculateTotalFare(numPassengers, destinationChoice, farePrices)
    fmt.Printf("Total fare for %d passengers: %.2f\n", numPassengers, totalFare)
}

// Function to calculate total fare for the journey
func calculateTotalFare(numPassengers int, route string, farePrices map[string]float64) float64 {
    fare := farePrices[route]
    return float64(numPassengers) * fare
}

// OUTPUT:
// Available Flights:
// Flight: FL001,  Departure: New York,    Destination: London,    Departure Time: 08:00,  Capacity: 200
// Flight: FL002,  Departure: London,      Destination: New York,  Departure Time: 09:00,  Capacity: 220
// Flight: FL003,  Departure: Tokyo,       Destination: Paris,     Departure Time: 10:00,  Capacity: 180

// Enter the flight number you want to book: FL001
// Enter the number of passengers: 3
// Enter the destination: London
// Enter name of passenger 1: jhon
// Enter name of passenger 2: peter
// Enter name of passenger 3: mark
// Total fare for 3 passengers: 1800.00