package restclient

import (
	"net/http"
	"testing"
)

// A good client
func goodClient() RestClient {
	return New("api.github.com")
}

// A bad client
func badClient() RestClient {
	return New("bad-api-github.doesNotExist")
}

func TestRestClient_Get(t *testing.T) {

	good := goodClient()
	bad := badClient()

	tests := []struct {
		name    string
		rc      *RestClient
		path    string
		want    responseWantData
		wantErr bool
	}{
		// Good test: query mit license from github
		{
			"Licenses Request github (good domain)", &good,
			"licenses/mit",
			responseWantData{
				http.StatusOK,
				http.MethodGet,
				"https://api.github.com/licenses/mit",
			},
			false, // No error
		},

		// Good test: query all licenses from github
		{
			"Licenses Request github (good domain)", &good,
			"licenses",
			responseWantData{
				http.StatusOK,
				http.MethodGet,
				"https://api.github.com/licenses",
			},
			false, // No error
		},

		// Bad test: query all licenses from github
		{
			"Licenses Request github (bad domain)", &bad,
			"licenses",
			responseWantData{
				0,
				http.MethodGet,
				"https://bad-api.github.com/licenses",
			},
			true, // Error cuz domain does not exist
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rc.Get(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestClient.Get() = %v, want %v", got, tt.want)
			}*/

			if tt.wantErr {
				t.Log("In this test we do not evaluete the repsonde, cuz there is none")
				return
			}

			if !tt.want.eq(got) {
				t.Errorf("RestClient.Get() = %v, want %v", got, tt.want)
			}

		})
	}
}

type responseWantData struct {
	status          int
	method, fullURL string
}

// A simple equal function
func (d *responseWantData) eq(r *http.Response) bool {
	return d.status == r.StatusCode &&
		d.method == r.Request.Method &&
		d.fullURL == r.Request.URL.String()
}
