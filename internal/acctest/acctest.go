package acctest

import (
	"testing"
	"time"
)

// These two functions come form the terraform-provider-aws code
// RunSerialTests1Level runs test cases in parallel, optionally sleeping between each.
func RunSerialTests1Level(t *testing.T, testCases map[string]func(t *testing.T), d time.Duration) {
	t.Helper()

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
			time.Sleep(d)
		})
	}
}

// RunSerialTests2Levels runs test cases in parallel, optionally sleeping between each.
func RunSerialTests2Levels(t *testing.T, testCases map[string]map[string]func(t *testing.T), d time.Duration) {
	t.Helper()

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			RunSerialTests1Level(t, m, d)
		})
	}
}
