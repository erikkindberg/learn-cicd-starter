package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got := GetAPIKey(nil)
	want := ""
}
