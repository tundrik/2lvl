package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		args        []string
		expected    string
		expectedErr bool
	}{
		{
			name:        "Select single field with TAB delimiter",
			input:       "1\t2\t3\n4\t5\t6\n",
			args:        []string{"-f", "2"},
			expected:    "2\n5\n",
			expectedErr: false,
		},
		{
			name:        "Select single field with different delimiter",
			input:       "1:2:3\n4:5:6\n",
			args:        []string{"-f", "2", "-d", ":"},
			expected:    "2\n5\n",
			expectedErr: false,
		},
		{
			name:        "Select multiple fields with TAB delimiter",
			input:       "1\t2\t3\n4\t5\t6\n",
			args:        []string{"-f", "1,3"},
			expected:    "1\t3\n4\t6\n",
			expectedErr: false,
		},
		{
			name:        "Select fields out of bounds",
			input:       "1\t2\t3\n4\t5\t6\n",
			args:        []string{"-f", "4"},
			expected:    "\n\n",
			expectedErr: false,
		},
		{
			name:        "Only process lines containing delimiter",
			input:       "1\t2\t3\n4\n5\t6\n",
			args:        []string{"-s"},
			expected:    "1\t2\t3\n5\t6\n",
			expectedErr: false,
		},
		{
			name:        "Error reading input",
			input:       "",
			args:        []string{},
			expected:    "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			output := new(bytes.Buffer)

			if test.input == "" {
				if output.String() != "" {
					t.Errorf("Expected empty output, Got output: %q", output.String())
				}
				return
			}

			err := cut(input, output, test.args...)
			if (err != nil) != test.expectedErr {
				t.Errorf("Expected error: %v, Got error: %v", test.expectedErr, err)
			}

			if output.String() != test.expected {
				t.Errorf("Expected output: %q, Got output: %q", test.expected, output.String())
			}
		})
	}
}
