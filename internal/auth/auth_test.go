package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headerKey string
		headerVal string
		queryKey  string
		queryVal  string
		expect    string
		expectErr string
	}{
		{
			name:      "no header or query",
			expectErr: "no authorization header or query param included",
		},
		{
			name:      "empty Authorization header",
			headerKey: "Authorization",
			expectErr: "no authorization header or query param included",
		},
		{
			name:      "malformed Authorization header",
			headerKey: "Authorization",
			headerVal: "-",
			expectErr: "malformed authorization header",
		},
		{
			name:      "wrong Authorization type",
			headerKey: "Authorization",
			headerVal: "Bearer xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			name:      "valid Authorization header",
			headerKey: "Authorization",
			headerVal: "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "",
		},
		{
			name:      "valid query param",
			queryKey:  "api_key",
			queryVal:  "query-key-123",
			expect:    "query-key-123",
			expectErr: "",
		},
		{
			name:      "both header and query, header takes priority",
			headerKey: "Authorization",
			headerVal: "ApiKey header-key-456",
			queryKey:  "api_key",
			queryVal:  "query-key-123",
			expect:    "header-key-456",
			expectErr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Создаём *http.Request
			req := &http.Request{
				Header: http.Header{},
				URL:    &url.URL{},
			}
			// Добавляем заголовок, если указан
			if tc.headerKey != "" {
				req.Header.Add(tc.headerKey, tc.headerVal)
			}
			// Добавляем query-параметр, если указан
			if tc.queryKey != "" {
				q := url.Values{}
				q.Add(tc.queryKey, tc.queryVal)
				req.URL.RawQuery = q.Encode()
			}

			got, err := GetAPIKey(req)
			if tc.expectErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.expectErr) {
					t.Errorf("Expected error containing %q, got %v", tc.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tc.expect {
				t.Errorf("Expected key %q, got %q", tc.expect, got)
			}
		})
	}
}
