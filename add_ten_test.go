package add_ten

import "testing"

func TestAddTen(t *testing.T) {
	x := 32
	y := add_ten(x)
	if y != 42 {
		t.Errorf("Expected %d got %d", 42, y)
	}
}
