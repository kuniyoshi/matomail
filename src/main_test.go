package main

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
	"testing"
)

func TestLineProcessor_ProcessLines(t *testing.T) {
	tests := []struct {
		name       string
		inputLines []string
		regex      string
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "basic functionality",
			inputLines: []string{"line1", "line2", "line2", "line2", "line3"},
			wantOutput: "line1\n(1) line2\n...\n(3) line2\nline3\n",
		},
		{
			name:       "regex pattern",
			inputLines: []string{"2025-05-14 10:00:00 ERROR Connection failed", "2025-05-14 10:01:00 ERROR Connection failed", "2025-05-14 10:02:00 ERROR Connection failed"},
			regex:      `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`,
			wantOutput: "(1) 2025-05-14 10:00:00 ERROR Connection failed\n...\n(3) 2025-05-14 10:02:00 ERROR Connection failed\n",
		},
		{
			name:       "empty input",
			inputLines: []string{},
			wantOutput: "",
		},
		{
			name:       "special characters",
			inputLines: []string{"line with !@#$%^&*()", "line with !@#$%^&*()", "line with different !@#$%^&*()"},
			wantOutput: "(1) line with !@#$%^&*()\n(2) line with !@#$%^&*()\nline with different !@#$%^&*()\n",
		},
		{
			name:       "invalid regex",
			inputLines: []string{"line1", "line2"},
			regex:      "[", // Invalid regex pattern
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			writer := bufio.NewWriter(&buf)
			scanner := bufio.NewScanner(strings.NewReader(strings.Join(tt.inputLines, "\n")))

			var re *regexp.Regexp
			var err error
			if tt.regex != "" {
				re, err = regexp.Compile(tt.regex)
				if err != nil {
					if !tt.wantErr {
						t.Errorf("Compile regex: %v", err)
					}
					return
				}
			}

			processor := newLineProcessor(writer, scanner, re)
			if err := processor.ProcessLines(); (err != nil) != tt.wantErr {
				t.Errorf("ProcessLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
