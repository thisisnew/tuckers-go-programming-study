package main

import "testing"

func TestSquare1(t *testing.T) {
	rst := square(9)

	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) return %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)

	if rst != 9 {
		t.Errorf("square(3) should be 9 but square(3) return %d", rst)
	}
}
