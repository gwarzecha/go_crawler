package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		rawBaseURL    string
		htmlBody      string
		expected      []string
		errorContains string
	}{
		{
			name:       "absolute and relative URLs",
			rawBaseURL: "https://blog.boot.dev",
			htmlBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected:      []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			errorContains: "",
		},
		{
			name:       "invalid HTML",
			rawBaseURL: "https://blog.boot.dev",
			htmlBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:       "invalid url",
			rawBaseURL: ":/blog.boot",
			htmlBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected:      nil,
			errorContains: "failed to parse base URL",
		},
		{
			name:       "multiple links with relative and absolute URLs",
			rawBaseURL: "https://blog.boot.dev",
			htmlBody: `
	<html>
			<body>
					<a href="/relative/path">Relative Link</a>
					<a href="https://blog.boot.dev/absolute/path">Absolute Link</a>
					<a href="relative/path/no/slash">Another Relative Link</a>
			</body>
	</html>
	`,
			expected: []string{
				"https://blog.boot.dev/relative/path",
				"https://blog.boot.dev/absolute/path",
				"https://blog.boot.dev/relative/path/no/slash",
			},
			errorContains: "",
		},
		{
			name:       "no links present",
			rawBaseURL: "https://blog.boot.dev",
			htmlBody: `
	<html>
			<body>
					<p>No links here</p>
			</body>
	</html>
	`,
			expected: nil,
		},
		{
			name:       "duplicate URLs",
			rawBaseURL: "https://example.com",
			htmlBody: `
	<html>
			<body>
					<a href="/page1">Page 1</a>
					<a href="/page2">Page 2</a>
					<a href="/page1">Page 1 Again</a>
			</body>
	</html>
	`,
			expected: []string{
				"https://example.com/page1",
				"https://example.com/page2",
				"https://example.com/page1",
			},
			errorContains: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.rawBaseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.htmlBody, baseURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("test %v - '%s' FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}

		})
	}
}
