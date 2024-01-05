package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, Go!",
		},
		{
			items:  []string{"Foo"},
			result: "Hello, Foo!",
		},
		{
			items:  []string{"Foo", "Bar"},
			result: "Hello, Foo, Bar!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s [%v], got %s", st.result, st.items, s)
		}
	}
}
