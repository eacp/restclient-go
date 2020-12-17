package restclient

import (
	"net/http"
)

// RestClient is a wrapper for an http client
// It also contains a host so all
// requests go to that post
type RestClient struct {
	baseURL    string
	httpClient *http.Client
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

// New returns a new rest client
// that uses the Default HTTP Client
// witha timeout of x seconds
// and configures the base url according
// to the host
func New(host string) RestClient {
	// Default settings
	return RestClient{
		makeBaseURL(false, host),
		http.DefaultClient,
	}
}

// NewLocalhost is a conveniance constructor
// that makes http calls to localhost using a specific
// port and http instead of https
func NewLocalhost(port string) (rc RestClient) {
	rc.httpClient = http.DefaultClient

	// If the user writes the ':' before the
	// port number then use only the number after
	if port[0] == ':' {
		rc.baseURL = makeBaseURL(true, "localhost"+port)
	} else {
		rc.baseURL = makeBaseURL(true, "localhost:"+port)
	}

	return
}

// NewInsecure returns a new rest client
// that uses the Default HTTP Client
// witha timeout of x seconds
// and configures the base url according
// to the host
// Uses HTTP instead of HTTPS
func NewInsecure(host string) RestClient {
	// Default settings
	return RestClient{
		makeBaseURL(true, host),
		http.DefaultClient,
	}
}
