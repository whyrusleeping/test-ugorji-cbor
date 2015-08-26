package canonicaltest

import (
	"testing"
	"testing/quick"
)

var testCases []interface{}

func init() {
	tc1 := map[string]string{
		"hello": "world",
	}

	tc2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	tc3 := map[string]interface{}{
		"a": 1,
		"b": "hello",
		"c": map[string]interface{}{
			"c/a": 1,
			"c/b": "world",
			"c/c": []int{1, 2, 3, 4},
		},
	}

	tc4 := map[string]interface{}{
		"a": 1,
		"b": "hello",
		"c": map[string]interface{}{
			"c/a": 1,
			"c/b": "world",
			"c/c": []int{1, 2, 3, 4},
			"c/d": map[string]interface{}{
				"c/d/a": "fdisajfoidsajfopdjsaopfjdsapofda",
				"c/d/b": "fdsafjdposakfodpsakfopdsakfpodsakfpodksaopfkdsopafkdopsa",
				"c/d/c": "poir02  ir30qif4p03qir0pogjfpoaerfgjp ofke[padfk[ewapf kdp[afep[aw",
				"c/d/d": "fdsopafkd[sa f-32qor-=4qeof -afo-erfo r-eafo 4e-  o r4-qwo ag",
				"c/d/e": "kfep[a sfkr0[paf[a foe-[wq  ewpfao-q ro3-q ro-4qof4-qor 3-e orfkropzjbvoisdb",
				"c/d/f": "",
			},
		},
	}

	testCases = []interface{}{tc1, tc2, tc3, tc4}
}

func TestRoundtripBasic(t *testing.T) {
	for _, tca := range testCases {
		var tcb map[string]interface{}
		RoundTripTest(t, &tca, &tcb)
	}
}

func TestRoundtripCheck(t *testing.T) {
	t.Skip("skipping TestRoundtripCheck")

	f := func(o1 map[string]map[string]string) bool {
		var o2 map[string]map[string]string
		return RoundTripTest(t, &o1, &o2)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
