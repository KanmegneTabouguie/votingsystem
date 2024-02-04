package utils

import "time"

// GetCurrentTimestamp returns the current timestamp as a string
func GetCurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
