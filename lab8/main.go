package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// WaterTransport represents a generic water transport.
type WaterTransport struct {
    Type     string `json:"type"`
    Name     string `json:"name"`
    Capacity int    `json:"capacity"`
}

// EncodeWaterTransport encodes a WaterTransport struct to JSON format.
func EncodeWaterTransport(transport WaterTransport) ([]byte, error) {
    return json.Marshal(transport)
}

// DecodeWaterTransport decodes JSON data into a WaterTransport struct.
func DecodeWaterTransport(data []byte) (WaterTransport, error) {
    var transport WaterTransport
    err := json.Unmarshal(data, &transport)
    return transport, err
}

func main() {
    var choice int
    var transport WaterTransport

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Add Water Transport")
        fmt.Println("2. Test JSON Encoding")
        fmt.Println("3. Test JSON Decoding")
        fmt.Println("4. Exit")
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Println("Enter Transport Type:")
            fmt.Scanln(&transport.Type)
            fmt.Println("Enter Transport Name:")
            fmt.Scanln(&transport.Name)
            fmt.Println("Enter Transport Capacity:")
            fmt.Scanln(&transport.Capacity)
        case 2:
            encoded, err := EncodeWaterTransport(transport)
            if err != nil {
                fmt.Println("Error encoding:", err)
            } else {
                fmt.Println("Encoded JSON:", string(encoded))
            }
        case 3:
            fmt.Println("Enter JSON data:")
            var jsonData string
            fmt.Scanln(&jsonData)
            decoded, err := DecodeWaterTransport([]byte(jsonData))
            if err != nil {
                fmt.Println("Error decoding:", err)
            } else {
                fmt.Println("Decoded Water Transport:")
                fmt.Printf("Type: %s\nName: %s\nCapacity: %d\n", decoded.Type, decoded.Name, decoded.Capacity)
            }
        case 4:
            fmt.Println("Exiting...")
            os.Exit(0)
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
