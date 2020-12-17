package restclient

import (
	"io"
	"net/http"
	"net/url"
)

// This file mostly tries to replicate the methods available in
// the http package, but using the base url instead

// Get is a wrapper function for http.Client Post function
// It uses the base url and the path to construct the post url
// TODO: Make sure authentication and tokens work
func (rc *RestClient) Get(path string) (*http.Response, error) {
	// gen url
	url := rc.makeURL(path)
	// make the request
	return rc.httpClient.Get(url)
}

// Post is a wrapper function for http.Client Post function
// It uses the base url and the path to construct the post url
// TODO: Make sure authentication and tokens work
func (rc *RestClient) Post(path, contentType string, body io.Reader) (*http.Response, error) {
	url := rc.makeURL(path)
	return rc.httpClient.Post(url, contentType, body)
}

// PostForm is a wrapper function for http.Client PostForm function
// It uses the base url and the path to construct the post url
// TODO: Make sure authentication and tokens work
func (rc *RestClient) PostForm(path string, data url.Values) (*http.Response, error) {
	// gen url
	url := rc.makeURL(path)
	// make the request
	return rc.httpClient.PostForm(url, data)
}
