package restclient

import (
	"net/http"
)

// HTTPClient is an interface with all the methods
// neccesary to make REST API requests
// It is meant to be an interface that is satisfied
// by http.Client, but opens the door for
// other http client implementations (such as testing for example)
type HTTPClient interface {
	Get(string) (*http.Response, error)
}

// RestClient is a wrapper for an http client
// It also contains a host so all
// requests go to that post
type RestClient struct {
	baseURL    string
	httpClient HTTPClient
	insecure   bool
}

// Create the base url witha host
func makeBaseURL(insecure bool, host string) string {
	var protocol string

	if insecure {
		protocol = "http://"
	} else {
		protocol = "https://"
	}

	return protocol + host + "/"

}

// Construct a URL from the base URL
func (rc *RestClient) makeURL(request string) string {
	return rc.baseURL + request
}
