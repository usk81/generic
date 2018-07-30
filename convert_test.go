package generic

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_asURL(t *testing.T) {
	u, _ := url.Parse(testURLString)

	tests := []struct {
		name        string
		args        interface{}
		wantResult  *url.URL
		wantIsValid ValidFlag
		wantErr     bool
	}{
		{
			name:        "nil",
			args:        nil,
			wantResult:  nil,
			wantIsValid: false,
			wantErr:     false,
		},
		{
			name:        "string",
			args:        testURLString,
			wantResult:  u,
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name:        "url.URL",
			args:        u,
			wantResult:  u,
			wantIsValid: true,
			wantErr:     false,
		},
		{
			name:        "int",
			args:        1000,
			wantResult:  nil,
			wantIsValid: false,
			wantErr:     true,
		},
		{
			name:        "bool",
			args:        true,
			wantResult:  nil,
			wantIsValid: false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotIsValid, err := asURL(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("asURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("asURL() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("asURL() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
