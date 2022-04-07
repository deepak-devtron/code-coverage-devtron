package code_coverage_devtron

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(1, 2), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}
