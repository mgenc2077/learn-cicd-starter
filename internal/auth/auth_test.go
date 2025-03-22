package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		"Authorization": {"ApiKey THE_KEY_HERE"},
		"Content-Type":  {"application/json"},
	}

	asd, err := GetAPIKey(headers)
	if (err != nil) || (asd != "THE_KEY_HERE") {
		t.Errorf("Got error %v", err)
	}

}
