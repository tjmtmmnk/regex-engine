package regexp

import "testing"

func TestMatch(t *testing.T) {
	re := NewRegexp("bana(na)*")
	expect := map[string]bool{
		"bana":   true,
		"banana": true,
		"banaNa": false,
		"apple":  false,
	}

	for k, v := range expect {
		isMatch := re.Match(k)
		if isMatch != v {
			t.Logf(k + " is not correct")
			t.Fail()
		}
	}
}
