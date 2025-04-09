package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type wantResult struct {
		string
		error
	}

	tests := map[string]struct {
		header http.Header
		want   wantResult
	}{
		"valid header key": {
			header: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			want: wantResult{
				string: "",
				error:  nil,
			},
		},

		"no header": {
			header: http.Header{},
			want: wantResult{
				string: "",
				error:  ErrNoAuthHeaderIncluded,
			},
		},

		"invalid apiKey": {
			header: http.Header{
				"Authorization": []string{"NotApiKey abc123"},
			},
			want: wantResult{
				string: "",
				error:  errors.New("malformed authorization header"),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotInfo, gotError := GetAPIKey(tc.header)

			got := wantResult{
				string: gotInfo,
				error:  gotError,
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
