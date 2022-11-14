package lib

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b

	if got := Add(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubstract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b

	if got := Substract(a, b); got != expected {
		t.Errorf("Substract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}
