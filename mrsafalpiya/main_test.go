package main

import (
	"reflect"
	"testing"

	"github.com/mrsafalpiya/get-sentry-blogs/scrapit"
)

func Test_getBlogs(t *testing.T) {
	type args struct {
		maxPostsPage   uint
		link           string
		blogClass      string
		blogInfoClass  string
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
				blogInfoClass:  BLOG_INFO_CLASS,
				blogLinkClass:  BLOG_LINK_CLASS,
				blogStyleClass: BLOG_STYLE_CLASS,
				styleAttrib:    STYLE_ATTRIB,
			},
			want: []scrapit.Blog{
				{
					Title:         "Application Profiling for Python",
					Info:          "Profiling is an important tool in every developer’s toolkit because it provides a granular view into the execution of your program from your…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/21/profiling-for-python/",
					Slug:          "profiling-for-python",
				},
				{
					Title:         "Application Profiling for Node.js",
					Info:          "Profiling is an important tool in every developer’s toolkit because it provides a granular view into the execution of your program from your…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/72Gv5QAj0Y58Qk0lZIzXgm/128eb409ae2e568787c3eb4f07b536cd/PerformanceLaunch-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/21/profiling-for-node-js/",
					Slug:          "profiling-for-node-js",
				},
				{
					Title:         "How we run our Python tests in hundreds of environments really fast",
					Info:          "Not in a reading mood? You also can watch the talk I gave at DjangoCon 2022. One of Sentries core company values is “for every developer…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/14/how-we-run-our-python-tests-in-hundreds-of-environments-really-fast/",
					Slug:          "how-we-run-our-python-tests-in-hundreds-of-environments-really-fast",
				},
				{
					Title:         "Python 3.11 Release - Top 5 Things to Know",
					Info:          "Python 3.11 was released on Oct. 24th, 2022. This latest version makes Python faster and even more user-friendly. If you’re not ready to…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/08/python-3-11-release-top-5-things-to-know/",
					Slug:          "python-3-11-release-top-5-things-to-know",
				},
				{
					Title:         "How Sentry uncovered an N+1 issue in djangoproject.com",
					Info:          "Sentry recently launched Performance Issues, a feature to help developers discover and fix common performance problems in their projects. We…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/ItvRLT7pjxn8KtckYxM0x/60a2a8675a07ce58cc8da3311b408645/metrics-meta__1_.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/04/how-sentry-uncovered-an-n-1-issue-in-djangoproject/",
					Slug:          "how-sentry-uncovered-an-n-1-issue-in-djangoproject",
				},
				{
					Title:         "A New Era of Sentry",
					Info:          "Today we are releasing Dynamic Sampling, available to all new customers, and opt-in for existing customers. This goes beyond a new feature…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7AiIutPpRCml9eKF2C8XjO/18413b1f1538c57a5998fbf020671981/DynamicSampling2_meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/11/02/a-new-era-of-sentry/",
					Slug:          "a-new-era-of-sentry",
				},
				{
					Title:         "We Just Gave $260,028 to Open Source Maintainers",
					Info:          "Sentry is an open source company, and it’s important to us to financially support our non-commercial colleagues in the community as we…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7GUQsSVuFa1gvDojQ2KvN7/0953c626442c0f18096e9bc7d90f3c9f/weOss-2022-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/27/we-just-gave-260-028-dollars-to-open-source-maintainers/",
					Slug:          "we-just-gave-260-028-dollars-to-open-source-maintainers",
				},
				{
					Title:         "Django Performance Improvements - Part 4: Caching in Django Applications",
					Info:          "In the first three parts of this series around improving performance in your Django applications, we focused on database, code optimization…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/b1Otcv6o3KZ3fqNQZoNLw/d64e4c5fa62cc5b7f79b8b4118a498e5/django-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/24/django-performance-improvements-part-4-caching-in-django-applications/",
					Slug:          "django-performance-improvements-part-4-caching-in-django-applications",
				},
				{
					Title:         "Top 3 Issue Alert Tips to Stop Noisy Notifications",
					Info:          "Sentry Alerts ping you on Slack, Microsoft Teams, or Pager Duty when something goes needs your attention. However, too many alerts can turn…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/7LLI0QOFlsyNpnThyXZyFr/51ec12e9014a31639ffb5d8f23641279/Errors3-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/20/top-3-issue-alert-tips-to-stop-noisy-notifications/",
					Slug:          "top-3-issue-alert-tips-to-stop-noisy-notifications",
				},
				{
					Title:         "Building a Performant iOS Profiler",
					Info:          "Here is a quick overview of profilers, and a deep dive into how we built the Sentry iOS profiler that has low enough overhead that it could run in production apps with minimal impact to user experience.",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/06/building-an-ios-profiler/",
					Slug:          "building-an-ios-profiler",
				},
				{
					Title:         "Spooky Season means Hacktoberfest",
					Info:          "🎃 Spooky Season means Hacktoberfest is here! 👻 Hacktoberfest is less spooky and more exciting for us here at Sentry. If you’re new to…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5xdgv1hdAKY4y7Uf2B1UrU/7e36eb5dc25ab53680e9d98022406f58/ecosystem3-meta__1_.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/05/spooky-season-means-hacktoberfest/",
					Slug:          "spooky-season-means-hacktoberfest",
				},
				{
					Title:         "Getting to That Elusive “Inbox Zero” With Custom Alerts and Codeowners",
					Info:          "\"I had to be able to balance my engineers’ time between fixing bugs and building new features, for that to happen we needed a solution that helped us stay on top of our backlog.\"",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/eIXOrYX5LjBuq03xunkPG/96757dab2600933d1b97455dd3636bb0/metric-alerts-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/04/inbox-zero-with-custom-alerts-and-codeowners/",
					Slug:          "inbox-zero-with-custom-alerts-and-codeowners",
				},
				{
					Title:         "Unity Performance Testing Tools & Benchmarks",
					Info:          "The following guest post addresses how to improve your services’ performance with Sentry and other application profilers for Unity. Learn…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
					PostLink:      "https://blog.sentry.io/2022/10/03/unity-performance-testing-tools-and-benchmarks/",
					Slug:          "unity-performance-testing-tools-and-benchmarks",
				},
				{
					Title:         "Python Performance Testing: A Comprehensive Guide",
					Info:          "The following guest post addresses how to improve your services’s performance with Sentry and other application profilers for Python. Check…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/71Q0i090p06pCDqQsj2aH2/34949d67c8f1020955219c8aa07bd26e/python-meta.jpg",
					PostLink:      "https://blog.sentry.io/2022/09/30/python-performance-testing-a-comprehensive-guide/",
					Slug:          "python-performance-testing-a-comprehensive-guide",
				},
				{
					Title:         "Unity Exception Handling: A Beginner’s Guide",
					Info:          "Exceptions are the outcomes you do not usually expect in your application. But as a developer, expecting the unexpected is essential to…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/45DqVWVynAMdGsB1keUoxt/6f315e73e1d086ccb06d9b2e84ad47df/unity-meta__current_.jpg",
					PostLink:      "https://blog.sentry.io/2022/09/30/unity-exception-handling-a-beginners-guide/",
					Slug:          "unity-exception-handling-a-beginners-guide",
				},
				{
					Title:         "Code-level Application Monitoring for Every Developer",
					Info:          "The monitoring, tooling, and observability space is crowded. It’s hard to keep track of what most tools in this category originally set out…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/1ml7JjcLxSFVLG4bhc5Hcd/1cbdb11fce30a44af4b9514304f9a322/Dex-Blog2-Meta-1200x630.png",
					PostLink:      "https://blog.sentry.io/2022/09/28/code-level-application-monitoring-for-every-developer/",
					Slug:          "code-level-application-monitoring-for-every-developer",
				},
				{
					Title:         "Deploy your Next.js application on Vercel using Sentry and GitHub Actions",
					Info:          "Thanks to the power of open source tooling and cloud services, shipping an application to production has never been that easy, In this blog…",
					ThumbnailLink: "https://images.ctfassets.net/em6l9zw4tzag/5oKTFLPvebgADv7SWHSMqA/4f282f6c579d1cc8667d8913d83a5668/meta-nextjs.png",
					PostLink:      "https://blog.sentry.io/2022/09/27/deploy-your-next-js-application-on-vercel-using-sentry-and-github-actions/",
					Slug:          "deploy-your-next-js-application-on-vercel-using-sentry-and-github-actions",
				},
				{
					Title:         "The Sentry Remix SDK is Now Available",
					Info:          "Sentry has made it a priority to support frontend JavaScript developers, regardless of the framework they use. This is why we have SDKs for…",
					ThumbnailLink: "https://blog.sentry.io/static/default-e58f27d48dbf46f8ddf19dc2404f62a9.png",
					PostLink:      "https://blog.sentry.io/2022/09/23/the-sentry-remix-sdk-is-now-available/",
					Slug:          "the-sentry-remix-sdk-is-now-available",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBlogs(tt.args.maxPostsPage, tt.args.link, tt.args.blogClass, tt.args.blogInfoClass, tt.args.blogLinkClass, tt.args.blogStyleClass, tt.args.styleAttrib)
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
