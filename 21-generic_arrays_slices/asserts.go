package generic_arrays_slices

import "testing"


func AssertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func AssertNotEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected == actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, actual bool) {
	t.Helper()

	if !actual {
		t.Errorf("expected true, got false")
	}
}

func AssertFalse(t *testing.T, actual bool) {
	t.Helper()

	if actual {
		t.Errorf("expected false, got true")
	}
}
