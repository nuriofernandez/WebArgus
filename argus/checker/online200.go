package checker

import (
	"fmt"
	"net/http"
)

func IsOnline(url string) bool {
	// http request
	fmt.Printf("[HTTP Checker] Requesting '%s'...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	// If document exists, will be 200 status code
	if resp.StatusCode == 200 {
		return true
	}

	return false
}
