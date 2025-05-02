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
				glob:   "**/api/v[0-9]",
				target: "http://example.com/api/v[0-9]",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/api/v[0-9]",
				target: "http://example.com/api/version",
			},
			want: false,
		},
		{
			args: args{
				glob:   "**/api\\?param",
				target: "http://example.com/api?param",
			},
			want: true,
		},
		{
			args: args{
				glob:   "**/api\\?param",
				target: "http://example.com/api-param",
			},
			want: false,
		},
		{
			args: args{
				glob:   "**/three-columns/settings.html\\?**id=settings-**",
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
	require.Equal(t, globMustToRegex("[a-z]").String(), `^\[a-z\]$`)
	require.Equal(t, globMustToRegex("$^+.\\*()|\\?\\{\\}\\[\\]").String(), `^\$\^\+\.\*\(\)\|\?\{\}\[\]$`)
}

func TestURLMatches(t *testing.T) {
	require.True(t, newURLMatcher("http://playwright.dev", nil).Matches("http://playwright.dev/"))
	require.True(t, newURLMatcher("http://playwright.dev?a=b", nil).Matches("http://playwright.dev/?a=b"))
	require.True(t, newURLMatcher("h*://playwright.dev", nil).Matches("http://playwright.dev/"))
	require.True(t, newURLMatcher("http://*.playwright.dev?x=y", nil).Matches("http://api.playwright.dev/?x=y"))
	require.True(t, newURLMatcher("**/foo/**", nil).Matches("http://playwright.dev/foo/bar"))
	require.True(t, newURLMatcher("?x=y", String("http://playwright.dev")).Matches("http://playwright.dev/?x=y"))
	require.True(t, newURLMatcher("./bar?x=y", String("http://playwright.dev/foo/")).Matches("http://playwright.dev/foo/bar?x=y"))

	// This is not supported, we treat ? as a query separator.
	require.False(t, globMustToRegex("http://localhost:8080/?imple/path.js").MatchString("http://localhost:8080/Simple/path.js"))
	require.False(t, newURLMatcher("http://playwright.?ev", nil).Matches("http://playwright.dev/"))
	require.True(t, newURLMatcher("http://playwright.?ev", nil).Matches("http://playwright./?ev"))
	require.False(t, newURLMatcher("http://playwright.dev/f??", nil).Matches("http://playwright.dev/foo"))
	require.True(t, newURLMatcher("http://playwright.dev/f??", nil).Matches("http://playwright.dev/f??"))
	require.True(t, newURLMatcher("http://playwright.dev\\?x=y", nil).Matches("http://playwright.dev/?x=y"))
	require.True(t, newURLMatcher("http://playwright.dev/\\?x=y", nil).Matches("http://playwright.dev/?x=y"))
	require.True(t, newURLMatcher("?bar", String("http://playwright.dev/foo")).Matches("http://playwright.dev/foo?bar"))
	require.True(t, newURLMatcher("\\\\?bar", String("http://playwright.dev/foo")).Matches("http://playwright.dev/foo?bar"))
	require.True(t, newURLMatcher("**/foo", String("http://first.host/")).Matches("http://second.host/foo"))
	require.True(t, newURLMatcher("*//localhost/", String("http://playwright.dev/")).Matches("http://localhost/"))
}
