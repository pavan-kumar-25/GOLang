package main

import (
	"testing"
)

func TestEncodeWaterTransport(t *testing.T) {
	transport := WaterTransport{
		Type:     "Ship",
		Name:     "Titanic",
		Capacity: 3000,
	}

	expected := `{"type":"Ship","name":"Titanic","capacity":3000}`
	encoded, err := EncodeWaterTransport(transport)
	if err != nil {
		t.Errorf("Error encoding: %v", err)
	}

	if string(encoded) != expected {
		t.Errorf("Encoding mismatch. Expected: %s, Got: %s", expected, string(encoded))
	}
}

func TestDecodeWaterTransport(t *testing.T) {
	data := []byte(`{"type":"Boat","name":"Speedy","capacity":10}`)

	expected := WaterTransport{
		Type:     "Boat",
		Name:     "Speedy",
		Capacity: 10,
	}

	decoded, err := DecodeWaterTransport(data)
	if err != nil {
		t.Errorf("Error decoding: %v", err)
	}

	if decoded != expected {
		t.Errorf("Decoding mismatch. Expected: %+v, Got: %+v", expected, decoded)
	}
}

// go mod init directory name
//go mod tidy
//go test
