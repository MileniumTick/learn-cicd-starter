package auth

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
  res, err := GetAPIKey(nil)

  if err == nil {
    t.Fatalf("expected error, got %v", err)
  }

  if res != "" {
    t.Fatal("expected empty response")
  }
}

