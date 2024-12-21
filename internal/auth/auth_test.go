package auth
import (
	"testing"
	"net/http"
)

func TestNoAuthHeaderErrors(t *testing.T) {
	noAuthHeader := &http.Header {}
	_, err := GetAPIKey(*noAuthHeader)
	// TODO(madhusudann): Fix condition after breaking test intentionally for CI
	if err ==  ErrNoAuthHeaderIncluded{
		t.Fatalf("Wanted error: %v but didn't receive one", ErrNoAuthHeaderIncluded)
	}
}

func TestValidApiKeyHeader(t *testing.T) {
        validHeader := &http.Header {}
	validHeader.Set("Authorization", "ApiKey abc")
        key, err := GetAPIKey(*validHeader)
        if err !=  nil || key != "abc" {
		t.Fatalf("Wanted key: %s, Got key: %s", "abc", key)
        }
}
