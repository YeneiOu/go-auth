package utils

import (
	"time"

	"github.com/google/uuid" // Import UUID package for generating UUIDs
)

// Trace generates a unique trace ID or identifier
func Trace() string {
	// Generate a UUID (Universal Unique Identifier) as a trace ID
	traceID := uuid.New().String()
	return traceID
}
func CallTimer(start int64) int64 {
	elapsed := time.Now().UnixMilli() - start
	return elapsed
}
