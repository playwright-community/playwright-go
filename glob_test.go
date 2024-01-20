package playwright

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_globMustToRegex(t *testing.T) {
	type args struct {
		glob   string
		target string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				glob:   "**/*.js",
				target: "https://localhost:8080/foo.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/*.css",
				target: "https://localhost:8080/foo.js",
			},
			want: false,
		},
		{
			args: args{
				glob:   "*.js",
				target: "https://localhost:8080/foo.js",
			},
			want: false,
		},
		{
			args: args{
				glob:   "https://**/*.js",
				target: "https://localhost:8080/foo.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "http://localhost:8080/simple/path.js",
				target: "http://localhost:8080/simple/path.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "http://localhost:8080/?imple/path.js",
				target: "http://localhost:8080/simple/path.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/{a,b}.js",
				target: "https://localhost:8080/a.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/{a,b}.js",
				target: "https://localhost:8080/b.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/{a,b}.js",
				target: "https://localhost:8080/c.js",
			},
			want: false,
		},
		{
			args: args{
				glob:   "**/*.{png,jpg,jpeg}",
				target: "https://localhost:8080/c.jpg",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/*.{png,jpg,jpeg}",
				target: "https://localhost:8080/c.jpeg",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/*.{png,jpg,jpeg}",
				target: "https://localhost:8080/c.png",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/*.{png,jpg,jpeg}",
				target: "https://localhost:8080/c.css",
			},
			want: false,
		},
		{
			args: args{
				glob:   "foo*",
				target: "foo.js",
			},
			want: true,
		},
		{
			args: args{
				glob:   "foo*",
				target: "foo/bar.js",
			},
			want: false,
		},
		{
			args: args{
				glob:   "http://localhost:3000/signin-oidc*",
				target: "http://localhost:3000/signin-oidc/foo",
			},
			want: false,
		},
		{
			args: args{
				glob:   "http://localhost:3000/signin-oidc*",
				target: "http://localhost:3000/signin-oidcnice",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/three-columns/settings.html?**id=[a-z]**",
				target: "http://mydomain:8080/blah/blah/three-columns/settings.html?id=settings-e3c58efe-02e9-44b0-97ac-dd138100cf7c&blah",
			},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("glob test %d", i), func(t *testing.T) {
			if got := globMustToRegex(tt.args.glob).MatchString(tt.args.target); got != tt.want {
				t.Errorf("globMustToRegex() = %v, want %v", got, tt.want)
			}
		})
	}

	require.Equal(t, globMustToRegex("\\?").String(), `^\?$`)
	require.Equal(t, globMustToRegex("\\").String(), `^\\$`)
	require.Equal(t, globMustToRegex("\\\\").String(), `^\\$`)
	require.Equal(t, globMustToRegex("\\[").String(), `^\[$`)
	require.Equal(t, globMustToRegex("[a-z]").String(), `^[a-z]$`)
	require.Equal(t, globMustToRegex("$^+.\\*()|\\?\\{\\}\\[\\]").String(), `^\$\^\+\.\*\(\)\|\?\{\}\[\]$`)
}
