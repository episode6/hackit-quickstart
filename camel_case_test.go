package main

import "testing"

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"whats-up-doc", "WhatsUpDoc"},
		{"whats_up_doc", "WhatsUpDoc"},
		{"whats up doc", "WhatsUpDoc"},
		{"whats up_doc-my man", "WhatsUpDocMyMan"},
		{"WhatsUpDoc", "WhatsUpDoc"},
	}

	for _, tst := range tests {
		output := camelCase(tst.input)
		if tst.expectedOutput != output {
			t.Errorf("camelCase(%v) = %v, but expected %v", tst.input, output, tst.expectedOutput)
		}
	}
}
