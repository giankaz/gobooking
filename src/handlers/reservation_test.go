package handlers

import (
	"testing"
)

func TestReservation(t *testing.T) {
	getTestCase(t, "/make-reservation")

	postTestCase(t, "/make-reservation", map[string]string{
		"first_name": "Test",
		"last_name":  "Test",
		"email":      "TESTE@Te.c",
		"phone":      "44444",
	})

}
