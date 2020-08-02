package eat

import "testing"

func TestEat(t *testing.T) {
	ret := GetResults()

	for i := range results {
		if ret == results[i] {
			return
		}
	}
	t.Fail()
}
