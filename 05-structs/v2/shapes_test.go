package structs

import "testing"

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}