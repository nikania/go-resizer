package util

import "testing"

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contains(arr, 1) {
		t.Errorf("Contains(arr, 1) = false, want true")
	}
	if Contains(arr, 6) {
		t.Errorf("Contains(arr, 6) = true, want false")
	}

	arr_strings := []string{"a", "b", "c", "d", "e"}
	if !Contains(arr_strings, "a") {
		t.Errorf("Contains(arr, \"a\") = false, want true")
	}
	if Contains(arr_strings, "f") {
		t.Errorf("Contains(arr, \"f\") = true, want false")
	}
}