package id

import (
	"fmt"

	"github.com/Bradkibs/TheGoTutsandPractices/machineid/distros"
)

func ID() (string, error) {

	id, err := distros.MachineID, distros.Error
	if err != nil && id == "" {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return id, nil
}

// ProtectedID returns a hashed version of the machine ID in a cryptographically secure way,
// using a fixed, application-specific key.
// Internally, this function calculates HMAC-SHA256 of the application ID, keyed by the machine ID.

func ProtectedID(appID string) (string, error) {
	id, err := ID()
	if err != nil {
		return "", fmt.Errorf("machineid Generation error: %v", err)
	}
	return distros.Protect(appID, id), nil
}
