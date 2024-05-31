package utils

import (
	"errors"

	"github.com/google/uuid"
)

// GenerateRandomID generates a random UUID (version 4) and returns it as a string.
//
// It returns the generated UUID as a string and an error if there was an issue generating the UUID.
// The error message includes the details of the original error.
//
// Returns:
// - uuid.UUID: The generated UUID.
// - error: An error if there was an issue generating the UUID, otherwise nil.
func GenerateRandomID() (uuid.UUID, error) {
	v4, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, errors.New("failed to generate random ID: " + err.Error())
	}

	return v4, nil
}
