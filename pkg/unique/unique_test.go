package unique

import "testing"

func TestString(t *testing.T) {
	var tests = []struct {
		Input    []string
		Expected []string
	}{
		{
			Input:    nil,
			Expected: nil,
		},
		{
			Input:    []string{},
			Expected: []string{},
		},
		{
			Input:    []string{"a", "b", "c", "d"},
			Expected: []string{"a", "b", "c", "d"},
		},
		{
			Input: []string{
				"a", "b", "c", "d", "a", "b", "c", "d",
			},
			Expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, test := range tests {
		res := Unique(test.Input)

		if len(res) != len(test.Expected) {
			t.Errorf(
				"expected %d values, but got %d",
				len(test.Expected), len(res),
			)
			return
		}
		for i := range test.Expected {
			if res[i] != test.Expected[i] {
				t.Errorf(
					"expected value %v but got %v",
					res[i], test.Expected[i],
				)
				return
			}
		}

		t.Logf("%v %v", res, test.Expected)
	}

}

func TestInt(t *testing.T) {
	var tests = []struct {
		Input    []int
		Expected []int
	}{
		{
			Input:    nil,
			Expected: nil,
		},
		{
			Input:    []int{},
			Expected: []int{},
		},
		{
			Input:    []int{1, 2, 3, 4},
			Expected: []int{1, 2, 3, 4},
		},
		{
			Input:    []int{1, 2, 3, 4, 1, 2, 3, 4},
			Expected: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		res := Unique(test.Input)

		if len(res) != len(test.Expected) {
			t.Errorf(
				"expected %d values, but got %d",
				len(test.Expected), len(res),
			)
			return
		}
		for i := range test.Expected {
			if res[i] != test.Expected[i] {
				t.Errorf(
					"expected value %v but got %v",
					res[i], test.Expected[i],
				)
				return
			}
		}

		t.Logf("%v %v", res, test.Expected)
	}

}
