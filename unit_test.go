package main

import (
  "testing"
)

func TestCalculateTransaction1 (t *testing.T){ // test transaction with $120.00
	points := calculateTransaction(120.00)
	if points != 90 {
		t.Errorf("calculateTransaction(120.00) = %d; want 90", points) 
	}
}

/* func TestFormatCheck1 (t *testing.T){
	bool := formatCheck(120.00)
	if bool != true {
		t.Errorf("formatCheck(120.00) = %t; want true", bool) 
	}
} */