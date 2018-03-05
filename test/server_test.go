package test

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := func(a int, b int) int { return a + b }(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFactorial(t *testing.T) {
	actual := func(a int) int {
		for i := a - 1; i > 0; i-- {
			a *= i
		}
		return a
	}(5)
	expected := 120
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestFailed(t *testing.T) {
	actual := 100
	expected := 120
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
