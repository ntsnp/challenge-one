package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mrsafalpiya/get-sentry-blogs/scrapit"
)

// Constants
// ---------

const LINK = "https://blog.sentry.io"
const BLOG_CLASS = ".css-spjg5j.e19gd7e57"
const BLOG_INFO_CLASS = ".css-1t38r8t.e19gd7e51"
const BLOG_LINK_CLASS = ".css-1qh9hqn.e19gd7e56"
const BLOG_STYLE_CLASS = ".e19gd7e53"
const STYLE_ATTRIB = "data-emotion"

// Arguments/Flags
// ---------------

var maxPostsPage *uint
var toPrintHelp *bool
var outputDir string

// Functions
// ---------

/*
 * Scrapes blogs from at most `maxPostsPage` post pages. Pass 0 to get all the
 * blogs.
 */
func getBlogs(maxPostsPage uint, link string, blogClass string, blogInfoClass string, blogLinkClass string, blogStyleClass string, styleAttrib string) ([]scrapit.Blog, error) {
	scrapitInstance, err := scrapit.NewScrapit(link)
	if err != nil {
		return scrapitInstance.Blogs, err
	}
	_ = scrapitInstance

	scrapitInstance.InitBlogsScrape(blogClass, blogInfoClass, blogLinkClass, blogStyleClass, styleAttrib)

	if maxPostsPage == 0 {
		log.Println("NOTE: All blog posts are begin scraped. This may take a while!")
	}

	err = scrapitInstance.Run(maxPostsPage, log.Default().Writer())
	if err != nil {
		return scrapitInstance.Blogs, err
	}

	return scrapitInstance.Blogs, nil
}

func usage(w io.Writer) {
	flag.CommandLine.SetOutput(w)
	fmt.Fprintf(w, "Usage: %s [options] output_dir\n\nWhere options are:\n", os.Args[0])
	flag.CommandLine.PrintDefaults()
	flag.CommandLine.Name()
}

func initFlags() {
	maxPostsPage = flag.Uint("p", 0, "Maximum number of posts page to scrape the blogs from")
	toPrintHelp = flag.Bool("help", false, "Print this help/usage message")

	flag.Parse()

	if *toPrintHelp {
		usage(os.Stdout)
		os.Exit(0)
	}

	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) != 1 {
		usage(os.Stderr)
		os.Exit(1)
	}
	outputDir = nonFlagArgs[0]
}

func main() {
	initFlags()

	blogs, err := getBlogs(*maxPostsPage, LINK, BLOG_CLASS, BLOG_INFO_CLASS, BLOG_LINK_CLASS, BLOG_STYLE_CLASS, STYLE_ATTRIB)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.Title)
		fmt.Println(blog.Info)
		fmt.Println(blog.ThumbnailLink)
		fmt.Println(blog.PostLink)
		fmt.Println(blog.Slug)
	}
}
