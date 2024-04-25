package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define a test case with a correct Authorization header
	testCases := []struct {
		name       string
		header     http.Header
		wantApiKey string
		wantErr    bool
	}{
		{
			name: "Correct Authorization Header",
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey 1234567890")
				return h
			}(),
			wantApiKey: "1234567890",
			wantErr:    false,
		},
		{
			name: "Missing Authorization Header",
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "")
				return h
			}(),
			wantApiKey: "1234567890",
			wantErr:    true,
		},
		{
			name: "Missing Authorization Header",
			header: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "1234567890")
				return h
			}(),
			wantApiKey: "1234567890",
			wantErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tc.header)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if gotApiKey != tc.wantApiKey {
					t.Errorf("Got %q, want %q", gotApiKey, tc.wantApiKey)
				}
			}
		})
	}
}
