package config

import (
	"os"
	"testing"
)

func Test_getEnvInt(t *testing.T) {
	tests := []struct {
		name           string
		variable       string
		def            int
		expectedOutput int
	}{
		{
			name:           "test happy path",
			variable:       "ENV_TEST",
			def:            10,
			expectedOutput: 100,
		},
		{
			name:           "test happy path - unknown",
			variable:       "UNKNOWN",
			def:            10,
			expectedOutput: 10,
		},
		{
			name:           "test happy path - empty",
			def:            10,
			expectedOutput: 10,
		},
		{
			name:           "test happy path - wrong type",
			variable:       "ENV_WRONG_TYPE",
			def:            10,
			expectedOutput: 10,
		},
	}
	for _, test := range tests {
		os.Setenv("ENV_TEST", "100")
		os.Setenv("ENV_WRONG_TYPE", "test")
		output := getEnvInt(test.variable, test.def)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, output)
		}
	}
}

func Test_getEnvString(t *testing.T) {
	tests := []struct {
		name           string
		variable       string
		def            string
		expectedOutput string
	}{
		{
			name:           "test happy path",
			variable:       "ENV_TEST",
			def:            "Local",
			expectedOutput: "Prod",
		},
		{
			name:           "test happy path - unknown",
			variable:       "UNKNOWN",
			def:            "Local",
			expectedOutput: "Local",
		},
		{
			name:           "test happy path - empty",
			def:            "Local",
			expectedOutput: "Local",
		},
	}
	for _, test := range tests {
		os.Setenv("ENV_TEST", "Prod")
		output := getEnvString(test.variable, test.def)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, output)
		}
	}
}
func Test_getEnvBool(t *testing.T) {
	tests := []struct {
		name           string
		variable       string
		def            bool
		expectedOutput bool
	}{
		{
			name:           "test happy path true",
			variable:       "ENV_TRUE",
			def:            false,
			expectedOutput: true,
		},
		{
			name:           "test happy path false",
			variable:       "ENV_FALSE",
			def:            true,
			expectedOutput: false,
		},
		{
			name:           "test happy path - unknown",
			variable:       "UNKNOWN",
			def:            true,
			expectedOutput: true,
		},
		{
			name:           "test happy path - empty",
			def:            true,
			expectedOutput: true,
		},
		{
			name:           "test happy path - wrong type",
			variable:       "ENV_WRONG_TYPE",
			def:            true,
			expectedOutput: true,
		},
	}
	for _, test := range tests {
		os.Setenv("ENV_TRUE", "true")
		os.Setenv("ENV_FALSE", "false")
		os.Setenv("ENV_WRONG_TYPE", "test")

		output := getEnvBool(test.variable, test.def)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, output)
		}
	}
}
