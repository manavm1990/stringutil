package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, ()_", "_)( ,olleH"},
		{"", ""}, 
	}
	for _, c := range cases { // Runs for loop 3x, in this case
		got := Reverse(c.in) // Keep resetting var got
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want) // %q gets replaced by each of the 3 parameters
		}
	}
}
