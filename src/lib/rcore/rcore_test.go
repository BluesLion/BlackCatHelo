package rcore

import "testing"

func TestRcore(t *testing.T) {
	retsults := []string{"A", "fgdf", "asdasdasd"}

	ret := PickOne(retsults)

	for i := range retsults {
		if ret == retsults[i] {
			return
		}
	}
	t.Fail()
}
