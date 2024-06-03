package pkg

import (
	"errors"

	"github.com/google/uuid"
)

// GenerateRandomID generates a random UUID (version 4) and returns it as a string.
//
// This function utilizes the uuid package to generate a UUID of version 4, which
// is based on random or pseudo-random numbers. If the UUID generation is successful,
// it returns the UUID. If there is an error during the generation process, it returns
// an empty UUID and an error containing the details of what went wrong.
//
// Returns:
// - uuid.UUID: The generated UUID.
// - error: An error if there was an issue generating the UUID, otherwise nil.
func GenerateRandomID() (uuid.UUID, error) {
	// Generate a new random (version 4) UUID.
	v4, err := uuid.NewRandom()
	if err != nil {
		// If there is an error during UUID generation, return an empty UUID and an error message.
		return uuid.UUID{}, errors.New("failed to generate random ID: " + err.Error())
	}

	// Return the generated UUID and no error.
	return v4, nil
}
