package main

import (
	"reflect"
	"testing"

	"github.com/challenge-one/scrapit"
)

func Test_getBlogs(t *testing.T) {
	type args struct {
		maxPostsPage   uint
		link           string
		blogClass      string
		blogLinkClass  string
		blogStyleClass string
		styleAttrib    string
	}
	tests := []struct {
		name    string
		args    args
		want    []scrapit.Blog
		wantErr bool
	}{
		{
			name: "Positive outcome",
			args: args{
				maxPostsPage:   1,
				link:           LINK,
				blogClass:      BLOG_CLASS,
				blogLinkClass:  BLOG_LINK_CLASS,
				blogStyleClass: BLOG_STYLE_CLASS,
				styleAttrib:    STYLE_ATTRIB,
			},
			want: []scrapit.Blog{
				{
					Title:         "Application Profiling for Python",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/21/profiling-for-python/",
				},
				{
					Title:         "Application Profiling for Node.js",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/72Gv5QAj0Y58Qk0lZIzXgm/128eb409ae2e568787c3eb4f07b536cd/PerformanceLaunch-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/21/profiling-for-node-js/",
				},
				{
					Title:         "How we run our Python tests in hundreds of environments really fast",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/14/how-we-run-our-python-tests-in-hundreds-of-environments-really-fast/",
				},
				{
					Title:         "Python 3.11 Release - Top 5 Things to Know",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/08/python-3-11-release-top-5-things-to-know/",
				},
				{
					Title:         "How Sentry uncovered an N+1 issue in djangoproject.com",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/ItvRLT7pjxn8KtckYxM0x/60a2a8675a07ce58cc8da3311b408645/metrics-meta__1_.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/04/how-sentry-uncovered-an-n-1-issue-in-djangoproject/",
				},
				{
					Title:         "A New Era of Sentry",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7AiIutPpRCml9eKF2C8XjO/18413b1f1538c57a5998fbf020671981/DynamicSampling2_meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/02/a-new-era-of-sentry/",
				},
				{
					Title:         "We Just Gave $260,028 to Open Source Maintainers",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7GUQsSVuFa1gvDojQ2KvN7/0953c626442c0f18096e9bc7d90f3c9f/weOss-2022-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/27/we-just-gave-260-028-dollars-to-open-source-maintainers/",
				},
				{
					Title:         "Django Performance Improvements - Part 4: Caching in Django Applications",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/b1Otcv6o3KZ3fqNQZoNLw/d64e4c5fa62cc5b7f79b8b4118a498e5/django-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/24/django-performance-improvements-part-4-caching-in-django-applications/",
				},
				{
					Title:         "Top 3 Issue Alert Tips to Stop Noisy Notifications",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7LLI0QOFlsyNpnThyXZyFr/51ec12e9014a31639ffb5d8f23641279/Errors3-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/20/top-3-issue-alert-tips-to-stop-noisy-notifications/",
				},
				{
					Title:         "Building a Performant iOS Profiler",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/06/building-an-ios-profiler/",
				},
				{
					Title:         "Spooky Season means Hacktoberfest",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5xdgv1hdAKY4y7Uf2B1UrU/7e36eb5dc25ab53680e9d98022406f58/ecosystem3-meta__1_.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/05/spooky-season-means-hacktoberfest/",
				},
				{
					Title:         "Getting to That Elusive “Inbox Zero” With Custom Alerts and Codeowners",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/eIXOrYX5LjBuq03xunkPG/96757dab2600933d1b97455dd3636bb0/metric-alerts-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/04/inbox-zero-with-custom-alerts-and-codeowners/",
				},
				{
					Title:         "Unity Performance Testing Tools & Benchmarks",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/03/unity-performance-testing-tools-and-benchmarks/",
				},
				{
					Title:         "Python Performance Testing: A Comprehensive Guide",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/09/30/python-performance-testing-a-comprehensive-guide/",
				},
				{
					Title:         "Unity Exception Handling: A Beginner’s Guide",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
					PostLink:      "https://blog.sentry.io/2022/09/30/unity-exception-handling-a-beginners-guide/",
				},
				{
					Title:         "Code-level Application Monitoring for Every Developer",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/1ml7JjcLxSFVLG4bhc5Hcd/1cbdb11fce30a44af4b9514304f9a322/Dex-Blog2-Meta-1200x630.png",
					PostLink:      "https://blog.sentry.io/2022/09/28/code-level-application-monitoring-for-every-developer/",
				},
				{
					Title:         "Deploy your Next.js application on Vercel using Sentry and GitHub Actions",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5oKTFLPvebgADv7SWHSMqA/4f282f6c579d1cc8667d8913d83a5668/meta-nextjs.png",
					PostLink:      "https://blog.sentry.io/2022/09/27/deploy-your-next-js-application-on-vercel-using-sentry-and-github-actions/",
				},
				{
					Title:         "The Sentry Remix SDK is Now Available",
					ThumbnailLink: "https://blog.sentry.io/static/default-e58f27d48dbf46f8ddf19dc2404f62a9.png",
					PostLink:      "https://blog.sentry.io/2022/09/23/the-sentry-remix-sdk-is-now-available/",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBlogs(tt.args.maxPostsPage, tt.args.link, tt.args.blogClass, tt.args.blogLinkClass, tt.args.blogStyleClass, tt.args.styleAttrib)
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
