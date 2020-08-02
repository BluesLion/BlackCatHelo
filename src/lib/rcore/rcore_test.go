package rcore

import "testing"

func TestRcore(t *testing.T) {
	retsults := []string{"A", "fgdf", "asdasdasd"}

	ret = PickOne(retsults)

	for i := range ret {
		if ret == ret[i] {
			return
		}
	}
	t.Fail()
}
