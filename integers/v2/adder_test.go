package integers

import (
	"fmt"
	"testing"
)

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

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleAddFloat() {
	sum := AddFloat(3.4, 6)
	fmt.Println(sum)
}
