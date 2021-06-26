package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestAddFloat(t *testing.T) {
	sum := AddFloat(3.4, 6)
	var expected float32 = 9.4

	if sum != expected {
		t.Errorf("expected '%f' but got '%f'", expected, sum)
	}
}
