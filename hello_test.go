package hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, Foo!"
	got := Say([]string{"Foo"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
