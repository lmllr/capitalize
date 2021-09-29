package capitalize

import (
	"testing"

	"github.com/lmllr/capitalize"
)

func TestStr(t *testing.T) {
	want := "Hello"
	if got := capitalize.Str("hello"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	want = "Hello"
	if got := capitalize.Str("HELLO"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	want = "123xyz"
	if got := capitalize.Str("123XYZ"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	want = "Foobar--"
	if got := capitalize.Str("foobar--"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	want = "Foobar--"
	if got := capitalize.Str("fooBAR--"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	want = "Foobar--"
	if got := capitalize.Str("FOOBAR--"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
