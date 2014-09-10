package piglatin

import "testing"

var translateTests = []struct {
	in  string
	out string
}{
	{
		in:  "dog",
		out: "ogday",
	},
	{
		in:  "cat",
		out: "atcay",
	},
	{
		in:  "hat",
		out: "athay",
	},
	{
		in:  "egg",
		out: "eggday",
	},
}

func TestTranslate(t *testing.T) {
	for i, test := range translateTests {
		actual := Translate(test.in)
		if actual != test.out {
			t.Errorf("%d failed", i)
		}
	}
}
