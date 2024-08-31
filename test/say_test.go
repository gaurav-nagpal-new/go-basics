package test

import "testing"

func TestSay(t *testing.T) {

	testCases := []struct {
		names  []string
		result string
	}{
		{
			result: "Hello, World",
		},
		{
			names:  []string{"Gaurav", "Simran"},
			result: "Hello, Gaurav Simran",
		},
	}

	for _, st := range testCases {
		if res := Say(st.names); st.result != res {
			t.Errorf("Wanted: %s Got: %s, Given: %v", st.result, res, st.names)
		}
	}
}
