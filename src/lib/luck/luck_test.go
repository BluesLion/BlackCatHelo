package luck

import "testing"

func TestLuck(t *testing.T) {
	ret := GetResults()

	for i := range results {
		if ret == results[i] {
			return
		}
	}
	t.Fail()
}
