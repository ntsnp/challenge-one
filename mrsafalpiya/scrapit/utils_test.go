package scrapit

import (
	"testing"
)

func Test_urlFromCSSVal(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive outcome #1",
			args: args{"url(foo)"},
			want: "foo",
		},
		{
			name: "Positive outcome #2 -- With leading slashes",
			args: args{"url(//foo)"},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlFromCSSVal(tt.args.v); got != tt.want {
				t.Errorf("urlFromCSSVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanUrl(t *testing.T) {
	type args struct {
		inputUrl string
		protocol string
		host     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive outcome #1",
			args: args{"http://foobar.com", "http://", "foobar.com"},
			want: "http://foobar.com",
		},
		{
			name: "Positive outcome #2",
			args: args{"https://foobar.com", "https://", "foobar.com"},
			want: "https://foobar.com",
		},
		{
			name: "Positive outcome #3",
			args: args{"foobar.com", "https://", "foobar.com"},
			want: "https://foobar.com",
		},
		{
			name: "Positive outcome #4",
			args: args{"search/1", "https://", "foobar.com"},
			want: "https://foobar.com/search/1",
		},
		{
			name: "Positive outcome #5",
			args: args{"/search/1", "https://", "foobar.com"},
			want: "https://foobar.com/search/1",
		},
		{
			name: "Positive outcome #6",
			args: args{"//search/1", "https://", "foobar.com"},
			want: "https://foobar.com/search/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanUrl(tt.args.inputUrl, tt.args.protocol, tt.args.host); got != tt.want {
				t.Errorf("cleanUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
