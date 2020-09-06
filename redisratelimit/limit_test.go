package redisratelimit

import (
	"testing"
)

func Test_getRemaining(t *testing.T) {
	tests := []struct {
		name           string
		maxQuotas      int
		used           int
		expectedOutput int
	}{
		{
			name:           "test happy path",
			maxQuotas:      5,
			used:           3,
			expectedOutput: 2,
		},
		{
			name:           "test happy exceeded",
			maxQuotas:      5,
			used:           6,
			expectedOutput: -1,
		},
	}
	for _, test := range tests {
		output := getRemaining(test.maxQuotas, test.used)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, output)
		}
	}
}
