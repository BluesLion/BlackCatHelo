package dndalign

import "testing"

func TestAlign(t *testing.T) {
	ret := GetResults()

	for i := range results {
		if ret == results[i] {
			return
		}
	}
	t.Fail()
}