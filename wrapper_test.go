package restclient

import (
	"net/http"
	"testing"
)

func Test_makeBaseURL(t *testing.T) {
	type args struct {
		host     string
		insecure bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Github",
			args{"api.github.com", false},
			"https://api.github.com/",
		},
		{
			"Github (Insecure)",
			args{"api.github.com", true},
			"http://api.github.com/",
		},
		{
			"Example",
			args{"api.example.com", false},
			"https://api.example.com/",
		},
		{
			"Example (Insecure)",
			args{"api.example.com", true},
			"http://api.example.com/",
		},
		{
			"Localhost",
			args{"localhost", true},
			"http://localhost/",
		},
		{
			"Localhost with a port",
			args{"localhost:8080", true},
			"http://localhost:8080/",
		},
		{
			"API in subdirectory",
			args{"example.com/api", false},
			"https://example.com/api/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeBaseURL(tt.args.insecure, tt.args.host); got != tt.want {
				t.Errorf("makeBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestClient_makeURL(t *testing.T) {

	base := "https://api.example.com/"
	c := RestClient{baseURL: base}

	tests := []struct {
		path, want string
	}{
		{"licenses", base + "licenses"},
		{"licenses/mit", base + "licenses/mit"},
		{"heroes/dc/batman", base + "heroes/dc/batman"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := c.makeURL(tt.path); got != tt.want {
				t.Errorf("RestClient.makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		host string
		want RestClient
	}{
		{
			"api.github.com",
			RestClient{
				"https://api.github.com/",
				http.DefaultClient,
			},
		},
		{
			"example.herokuapp.com",
			RestClient{
				"https://example.herokuapp.com/",
				http.DefaultClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			got := New(tt.host)
			want := tt.want
			if got.httpClient != want.httpClient || got.baseURL != want.baseURL {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
