/*

Scenario: Building a Package Delivery System

You have been tasked with developing a simple package delivery system in Go for a local courier service.
The system should allow users to schedule package deliveries, track the status of their packages,
and calculate delivery fees based on various factors. Your goal is to implement different types of functions
and error handling methods to ensure the reliability and usability of the system.

Package Management:
Define a struct named Package to represent each package, including attributes such as
PackageID, Sender, Recipient, Destination, Weight, and DeliveryStatus.

Delivery Scheduling:
Implement functions to schedule package deliveries, update delivery status, and track package movements.
Use a slice or map to store the details of scheduled packages.

Fee Calculation:
Develop a function called calculateDeliveryFee that takes package weight, distance, and urgency as input parameters and returns the total delivery fee.
Handle different fee calculation scenarios based on weight, distance, and delivery urgency.

Error Handling:
Implement error handling mechanisms to handle scenarios such as invalid package details, duplicate package IDs, and calculation errors.
Use custom error types to provide meaningful error messages to users and developers.

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Shippable interface for packages that can be shipped
type Shippable interface {
	getWeight() float64 // Method to get the weight of the package
}

// Package struct represents a package with its details
type Package struct {
	Shippable
	PackageID     int
	Sender        string
	Recipient     string
	Destination   string
	Weight        float64
	DeliveryStatus string
}

// SmallPackage struct represents a small package
type SmallPackage struct {
	Package
}

// LargePackage struct represents a large package
type LargePackage struct {
	Package
}

// getWeight method for SmallPackage
func (sp *SmallPackage) getWeight() float64 {
	return sp.Weight
}

// getWeight method for LargePackage
func (lp *LargePackage) getWeight() float64 {
	return lp.Weight
}

var scheduledPackages []Shippable // Slice to store scheduled packages

func main() {
	// Main function to manage the package delivery system

	for {
		// Continuous loop for user interaction

		// Display menu options
		fmt.Print("\nWelcome to the Package Delivery System!\n")
		fmt.Print("1. Schedule a package delivery\n")
		fmt.Print("2. Update delivery status\n")
		fmt.Print("3. Track package movements\n")
		fmt.Print("4. Calculate delivery fee\n")
		fmt.Print("5. Exit\n")

		var choice int
		fmt.Print("Please enter your choice (1-5):\n")
		fmt.Scanln(&choice)

		// Process user choice
		switch choice {
		case 1:
			schedule() // Schedule a new package delivery
		case 2:
			updateDeliveryStatus() // Update delivery status of a package
		case 3:
			trackPackageMovements() // Track movements of a package
		case 4:
			calculateDeliveryFeeWrapper() // Calculate delivery fee for a package
		case 5:
			fmt.Print("Exiting...") // Exit the program
			return
		default:
			fmt.Print("Invalid input..") // Handle invalid input
			break
		}
	}
}

// Function to schedule a new package delivery
func schedule() {
	var newPackage Shippable

	// Prompt user to input package details
	fmt.Print("Enter Package ID: ")
	var packageIDInput string
	fmt.Scanln(&packageIDInput)

	packageID, err := strconv.Atoi(packageIDInput)
	if err != nil {
		fmt.Println("Invalid Package ID. Please enter a valid integer.")
		return
	}

	var packageType string
	fmt.Print("Enter Package Type (small/large): ")
	fmt.Scanln(&packageType)

	fmt.Print("Enter Sender: ")
	var sender string
	fmt.Scanln(&sender)

	fmt.Print("Enter Recipient: ")
	var recipient string
	fmt.Scanln(&recipient)

	fmt.Print("Enter Destination: ")
	var destination string
	fmt.Scanln(&destination)

	fmt.Print("Enter Weight: ")
	var weight float64
	fmt.Scanln(&weight)

	if packageType == "small" {
		newPackage = &SmallPackage{Package{
			PackageID:     packageID,
			Sender:        sender,
			Recipient:     recipient,
			Destination:   destination,
			Weight:        weight,
			DeliveryStatus: "Scheduled",
		}}
	} else if packageType == "large" {
		newPackage = &LargePackage{Package{
			PackageID:     packageID,
			Sender:        sender,
			Recipient:     recipient,
			Destination:   destination,
			Weight:        weight,
			DeliveryStatus: "Scheduled",
		}}
	} else {
		fmt.Println("Invalid package type. Please enter 'small' or 'large'.")
		return
	}

	scheduledPackages = append(scheduledPackages, newPackage) // Add package to scheduled list

	fmt.Println("Package scheduled successfully!")
}

// Function to update delivery status of a package
func updateDeliveryStatus() {
	var packageIDInput string

	// Prompt user to input package ID
	fmt.Print("Enter Package ID: ")
	fmt.Scanln(&packageIDInput)

	packageID, err := strconv.Atoi(packageIDInput)
	if err != nil {
		fmt.Println("Invalid Package ID. Please enter a valid integer.")
		return
	}

	var newStatus string

	fmt.Print("Enter New Delivery Status: ")
	fmt.Scanln(&newStatus)

	// Iterate through scheduled packages to find the package by ID
	for _, pkg := range scheduledPackages {
		switch p := pkg.(type) {
		case *SmallPackage:
			if p.PackageID == packageID {
				p.DeliveryStatus = newStatus
				fmt.Println("Delivery status updated successfully!")
				return
			}
		case *LargePackage:
			if p.PackageID == packageID {
				p.DeliveryStatus = newStatus
				fmt.Println("Delivery status updated successfully!")
				return
			}
		}
	}

	fmt.Println("Package not found or delivery status update failed.")
}

// Function to track movements of a package
func trackPackageMovements() {
	var packageIDInput string

	// Prompt user to input package ID
	fmt.Print("Enter Package ID: ")
	fmt.Scanln(&packageIDInput)

	packageID, err := strconv.Atoi(packageIDInput)
	if err != nil {
		fmt.Println("Invalid Package ID. Please enter a valid integer.")
		return
	}

	// Iterate through scheduled packages to find the package by ID
	for _, pkg := range scheduledPackages {
		switch p := pkg.(type) {
		case *SmallPackage:
			if p.PackageID == packageID {
				// Display package details
				fmt.Printf("Package ID: %d\nSender: %s\nRecipient: %s\nDestination: %s\nDelivery Status: %s\n",
					p.PackageID, p.Sender, p.Recipient, p.Destination, p.DeliveryStatus)
				return
			}
		case *LargePackage:
			if p.PackageID == packageID {
				// Display package details
				fmt.Printf("Package ID: %d\nSender: %s\nRecipient: %s\nDestination: %s\nDelivery Status: %s\n",
					p.PackageID, p.Sender, p.Recipient, p.Destination, p.DeliveryStatus)
				return
			}
		}
	}

	fmt.Println("Package not found.")
}

// Function to interactively calculate delivery fee for a package
func calculateDeliveryFeeWrapper() {
	var weight, distance float64
	var urgency string

	// Prompt user to input weight
	fmt.Print("Enter Weight: ")
	_, err := fmt.Scanln(&weight)
	if err != nil {
		fmt.Println("Invalid input for Weight. Please enter a number.")
		return
	}

	// Prompt user to input distance
	fmt.Print("Enter Distance: ")
	_, err = fmt.Scanln(&distance)
	if err != nil {
		fmt.Println("Invalid input for Distance. Please enter a number.")
		return
	}

	fmt.Print("Enter Urgency (express/standard): ")
	_, err = fmt.Scanln(&urgency)
	if err != nil {
		fmt.Println("Invalid input for Urgency. Please enter 'express' or 'standard'.")
		return
	}

	// Convert input to lowercase for case-insensitive comparison
	urgency = strings.ToLower(urgency)

	// Validate urgency input
	if urgency != "express" && urgency != "standard" {
		fmt.Println("Invalid input for Urgency. Please enter 'express' or 'standard'.")
		return
	}

	// Call the calculateDeliveryFee function and display the result
	fee := calculateDeliveryFee(weight, distance, urgency)
	fmt.Println("Delivery fee:", fee)
}

// Function to calculate delivery fee based on weight, distance, and urgency
func calculateDeliveryFee(weight float64, distance float64, urgency string) float64 {
	// Calculate delivery fee based on weight, distance, and urgency

	baseFee := 10.0         // Base fee for delivery
	weightFee := weight * 2 // Fee based on weight

	var distanceFee float64
	switch {
	case distance < 100:
		distanceFee = 5
	case distance < 500:
		distanceFee = 10
	default:
		distanceFee = 15
	}

	var urgencyFee float64
	if urgency == "express" {
		urgencyFee = 15
	} else {
		urgencyFee = 5
	}

	// Calculate total delivery fee
	deliveryFee := baseFee + weightFee + distanceFee + urgencyFee

	return deliveryFee
}
