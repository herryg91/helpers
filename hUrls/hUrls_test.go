package hUrls

import (
	"testing"
)

func TestIsURL(t *testing.T) {
	type args struct {
		urlString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OK http",
			args: args{urlString: "http://golang.org"},
			want: true,
		},
		{
			name: "OK https",
			args: args{urlString: "https://golang.org"},
			want: true,
		},
		{
			name: "Wrong URL",
			args: args{urlString: "http:golang.org"},
			want: false,
		},
		{
			name: "Wrong Scheme",
			args: args{urlString: "htt://golang.org"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.urlString); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecureURL(t *testing.T) {
	type args struct {
		urlString string
	}
	tests := []struct {
		name          string
		args          args
		wantSecureUrl string
	}{
		{
			name:          "OK",
			args:          args{urlString: "http://golang.org"},
			wantSecureUrl: "https://golang.org",
		},
		{
			name:          "Wrong URL",
			args:          args{urlString: "http:golang.org"},
			wantSecureUrl: "http:golang.org",
		},
		{
			name:          "HTTP as params",
			args:          args{urlString: "http://golang.org?redirect=http://golang.com"},
			wantSecureUrl: "https://golang.org?redirect=http://golang.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSecureUrl := SecureURL(tt.args.urlString); gotSecureUrl != tt.wantSecureUrl {
				t.Errorf("SecureURL() = %v, want %v", gotSecureUrl, tt.wantSecureUrl)
			}
		})
	}
}
