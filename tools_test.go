package toolkit

import "testing"

func Test_Tools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)

	if len(s) != 10 {
		t.Error("wrong length of random string string returned")
	}
}
