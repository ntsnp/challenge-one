package main

import (
	"reflect"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanUrl(tt.args.inputUrl, tt.args.protocol, tt.args.host); got != tt.want {
				t.Errorf("cleanUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBlogs(t *testing.T) {
	type args struct {
		link       string
		divClass   string
		bgDivClass string
	}
	tests := []struct {
		name    string
		args    args
		want    []blog
		wantErr bool
	}{
		{
			name: "Positive outcome",
			args: args{
				link:       LINK,
				divClass:   BLOG_DIV_CLASS,
				bgDivClass: BLOG_BG_DIV_CLASS,
			},
			want: []blog{
				{
					title:         "Application Profiling for Python",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
				},
				{
					title:         "Application Profiling for Node.js",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/72Gv5QAj0Y58Qk0lZIzXgm/128eb409ae2e568787c3eb4f07b536cd/PerformanceLaunch-meta.jpg",
				},
				{
					title:         "How we run our Python tests in hundreds of environments really fast",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
				},
				{
					title:         "Python 3.11 Release - Top 5 Things to Know",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
				},
				{
					title:         "How Sentry uncovered an N+1 issue in djangoproject.com",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/ItvRLT7pjxn8KtckYxM0x/60a2a8675a07ce58cc8da3311b408645/metrics-meta__1_.jpg",
				},
				{
					title:         "A New Era of Sentry",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7AiIutPpRCml9eKF2C8XjO/18413b1f1538c57a5998fbf020671981/DynamicSampling2_meta.jpg",
				},
				{
					title:         "We Just Gave $260,028 to Open Source Maintainers",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7GUQsSVuFa1gvDojQ2KvN7/0953c626442c0f18096e9bc7d90f3c9f/weOss-2022-meta.jpg",
				},
				{
					title:         "Django Performance Improvements - Part 4: Caching in Django Applications",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/b1Otcv6o3KZ3fqNQZoNLw/d64e4c5fa62cc5b7f79b8b4118a498e5/django-meta.jpg",
				},
				{
					title:         "Top 3 Issue Alert Tips to Stop Noisy Notifications",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7LLI0QOFlsyNpnThyXZyFr/51ec12e9014a31639ffb5d8f23641279/Errors3-meta.jpg",
				},
				{
					title:         "Building a Performant iOS Profiler",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
				},
				{
					title:         "Spooky Season means Hacktoberfest",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5xdgv1hdAKY4y7Uf2B1UrU/7e36eb5dc25ab53680e9d98022406f58/ecosystem3-meta__1_.jpg",
				},
				{
					title:         "Getting to That Elusive “Inbox Zero” With Custom Alerts and Codeowners",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/eIXOrYX5LjBuq03xunkPG/96757dab2600933d1b97455dd3636bb0/metric-alerts-meta.jpg",
				},
				{
					title:         "Unity Performance Testing Tools & Benchmarks",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
				},
				{
					title:         "Python Performance Testing: A Comprehensive Guide",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
				},
				{
					title:         "Unity Exception Handling: A Beginner’s Guide",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
				},
				{
					title:         "Code-level Application Monitoring for Every Developer",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/1ml7JjcLxSFVLG4bhc5Hcd/1cbdb11fce30a44af4b9514304f9a322/Dex-Blog2-Meta-1200x630.png",
				},
				{
					title:         "Deploy your Next.js application on Vercel using Sentry and GitHub Actions",
					thumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5oKTFLPvebgADv7SWHSMqA/4f282f6c579d1cc8667d8913d83a5668/meta-nextjs.png",
				},
				{
					title:         "The Sentry Remix SDK is Now Available",
					thumbnailLink: "https://blog.sentry.io/static/default-e58f27d48dbf46f8ddf19dc2404f62a9.png",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBlogs(tt.args.link, tt.args.divClass, tt.args.bgDivClass)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBlogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBlogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
