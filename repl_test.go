package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "This is a test",
			expected: []string{
				"this",
				"is",
				"a",
				"test",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Incorrect length: %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWOrd := actual[i]
			expectedWord := cs.expected[i]
			if actualWOrd != expectedWord {
				t.Errorf("%v", actualWOrd, expectedWord)
			}
		}
	}

}
