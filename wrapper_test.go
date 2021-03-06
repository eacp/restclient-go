package restclient

import (
	"net/http"
	"testing"
	"time"
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

func TestNewLocalhost(t *testing.T) {
	tests := []struct {
		name, port string
		wantRc     RestClient
	}{
		{
			"Test with ':'", ":8080",
			RestClient{
				"http://localhost:8080/",
				http.DefaultClient,
			},
		},
		{
			"Test without ':'", "8080",
			RestClient{
				"http://localhost:8080/",
				http.DefaultClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewLocalhost(tt.port)
			r := tt.wantRc
			if g.baseURL != r.baseURL || g.httpClient != r.httpClient {
				t.Errorf("NewLocalhost() = %v, want %v", g, tt.wantRc)
			}
		})
	}
}

func TestNewInsecure(t *testing.T) {
	tests := []struct {
		host string
		want RestClient
	}{
		{
			"api.github.com",
			RestClient{
				"http://api.github.com/",
				http.DefaultClient,
			},
		},
		{
			"example.herokuapp.com",
			RestClient{
				"http://example.herokuapp.com/",
				http.DefaultClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			got := NewInsecure(tt.host)
			want := tt.want
			if got.httpClient != want.httpClient || got.baseURL != want.baseURL {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithClient(t *testing.T) {
	// Create a custom http client here
	customClient := &http.Client{
		Timeout: time.Minute,
	}

	tests := []struct {
		host string
		want RestClient
	}{
		{
			"api.github.com",
			RestClient{
				"https://api.github.com/",
				customClient,
			},
		},
		{
			"example.herokuapp.com",
			RestClient{
				"https://example.herokuapp.com/",
				customClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			got := NewWithClient(tt.host, customClient)
			want := tt.want
			if got.httpClient != want.httpClient || got.baseURL != want.baseURL {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithClientInsecure(t *testing.T) {
	// Create a custom http client here
	customClient := &http.Client{
		Timeout: time.Minute,
	}

	tests := []struct {
		host string
		want RestClient
	}{
		{
			"api.github.com",
			RestClient{
				"http://api.github.com/",
				customClient,
			},
		},
		{
			"example.herokuapp.com",
			RestClient{
				"http://example.herokuapp.com/",
				customClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			got := NewWithClientInsecure(tt.host, customClient)
			want := tt.want
			if got.httpClient != want.httpClient || got.baseURL != want.baseURL {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithTimeout(t *testing.T) {
	// We will only test if the internal http client
	// has the correct timeout

	type args struct {
		host    string
		timeout time.Duration
	}
	tests := []args{
		{"api.github.com", time.Minute},
		{"slow-api.example.com", time.Minute * 5},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			/*if got := NewWithTimeout(tt.args.host, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithTimeout() = %v, want %v", got, tt.want)
			}*/

			got := NewWithTimeout(tt.host, tt.timeout)
			hc := got.httpClient

			if hc.Timeout != tt.timeout {
				t.Errorf("NewWithTimeout() = %v, want %v", got, tt.timeout)
			}

		})
	}
}
