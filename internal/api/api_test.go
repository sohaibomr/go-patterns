package api

import "testing"

func TestSum(t *testing.T) {

	v := Sum(2, 2)
	if v != 4 {
		t.Error("Test Fails, expected 4")
	}
}
