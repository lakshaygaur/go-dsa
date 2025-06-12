package fibonaccinumber

import "testing"

func TestFibonacciNumber(t *testing.T) {

	t.Run("2", func(t *testing.T) {
		got := fib(2)
		want := 1

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("3", func(t *testing.T) {
		got := fib(3)
		want := 2

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("4", func(t *testing.T) {
		got := fib(4)
		want := 3

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
