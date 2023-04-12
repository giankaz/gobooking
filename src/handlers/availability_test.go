package handlers

import (
	"testing"
)

func TestAvailability(t *testing.T) {
	getTestCase(t, "/search-availability")

	postTestCase(t, "/search-availability", map[string]string{
		"start": "2025-10-10",
		"end":   "2026-10-10",
	})
}
