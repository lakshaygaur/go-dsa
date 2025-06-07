package lcstringcompression

import "testing"

func TestCompress(t *testing.T) {

	t.Run("['a', 'a', 'b', 'b', 'c', 'c', 'c']", func(t *testing.T) {
		in := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
		got := compress(in)
		want := 6
		out_str := "a2b2c3"

		if got != want && out_str != string(in[:got]) {
			t.Errorf("got %q, wanted %q && expected %q, received %q", got, want, out_str, string(in[:got]))
		}
	})
	// []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	t.Run("['a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b']", func(t *testing.T) {
		in := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
		got := compress(in)
		want := 4
		out_str := "ab12"

		if got != want && out_str != string(in[:got]) {
			t.Errorf("got %q, wanted %q && expected %q, received %q", got, want, out_str, string(in[:got]))
		}
	})
}
