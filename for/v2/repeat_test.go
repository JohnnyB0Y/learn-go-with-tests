package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatGo(t *testing.T) {
	repeated := RepeatGo("b")
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
