package main

import (
	"testing"
)

func TestFunction(t *testing.T) {
	t.Run("testing for int ", func(t *testing.T) {
		got := Addition(23, 27)
		want := 50
		checking(t, want, got)
	})
	t.Run("testing for string ", func(t *testing.T) {
		got := Addition("first_name ", "second_name")
		want := "first_name second_name"
		checking(t, want, got)
	})
	t.Run("testing for float32 ", func(t *testing.T) {
		got := Addition(32.5, 31.5)
		want := 64.0
		checking(t, want, got)
	})
}

func checking(t testing.TB, want, got any) {
	t.Helper()
	if want != got {
		t.Errorf("[Want] %q [Got] %q", want, got)
	}
}
