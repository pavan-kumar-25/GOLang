/*Scenario:
A logistics company named "FastCargo" needs to dispatch several delivery trucks to 
different destinations for transporting goods. Each truck has a unique ID and is 
assigned a specific route. The trucks need to travel concurrently to their destinations.*/

package main

import (
    "fmt"
    "sync"
)

// Vehicle represents a delivery truck with its ID, route, and status.
type Vehicle struct {
    ID     int
    Route  string
    Status string
}

// travel simulates the travel of a delivery truck on its route.
func travel(vehicle Vehicle, wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()

    // Simulating travel time
    fmt.Printf("Truck %d on route %s is traveling...\n", vehicle.ID, vehicle.Route)

    // Updating status
    vehicle.Status = "Traveling(updated status)"

    // Sending status update to channel
    ch <- fmt.Sprintf("Truck %d is %s", vehicle.ID, vehicle.Status)
}

func main() {
    var wg sync.WaitGroup

    // Channel to receive status updates
    statusChan := make(chan string)

    // Number of trucks
    var numTrucks int
    fmt.Println("Welcome to FastCargo!")
    fmt.Println("Enter the number of delivery trucks:")
    _, err := fmt.Scan(&numTrucks)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Input truck data
    trucks := make([]Vehicle, numTrucks)
    for i := 0; i < numTrucks; i++ {
        var route string
        fmt.Printf("Enter route for Truck %d: ", i+1)
        _, err := fmt.Scan(&route)
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }
        trucks[i] = Vehicle{ID: i + 1, Route: route, Status: "Idle"}
    }

    // Start goroutines for each truck
    for _, truck := range trucks {
        wg.Add(1)
        go travel(truck, &wg, statusChan)
    }

    // Wait for all goroutines to finish
    go func() {
        wg.Wait()
        close(statusChan)
    }()

    // Receive and print status updates from channel
    fmt.Println("\nCurrent Status:")
    for status := range statusChan {
        fmt.Println(status)
    }
    fmt.Println("All trucks have reached their destinations. Thank you for using FastCargo!")
}
