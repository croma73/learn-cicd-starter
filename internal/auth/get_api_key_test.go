package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("expected an error for missing header, but got nil")
	}

}

func TestGetAPIKey_HappyPath(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey this-is-the-secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := "this-is-the-secret-key"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
