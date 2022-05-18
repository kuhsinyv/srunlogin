package probe

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	// ErrNetworkLoggedIn the network was logged in
	ErrNetworkLoggedIn = errors.New("已登录网络")
)

// Probe the authentication home page of the connected network
func Probe(testURLString string) (string, error) {
	testURL, err := url.Parse(testURLString)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(testURL.String())
	if err != nil {
		return "", err
	}

	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	loginURL := resp.Request.URL
	if strings.EqualFold(loginURL.Host, testURL.Host) {
		return "", ErrNetworkLoggedIn
	}

	return loginURL.String(), nil
}
