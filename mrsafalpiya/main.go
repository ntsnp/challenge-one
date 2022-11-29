package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

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
var toSave *bool
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
		log.Println("[NOTE] All blog posts are begin scraped. This may take a while!")
	}

	err = scrapitInstance.Run(maxPostsPage, log.Default().Writer())
	if err != nil {
		return scrapitInstance.Blogs, err
	}

	return scrapitInstance.Blogs, nil
}

func saveBlogs(blogs []scrapit.Blog, outputDir string) {
	log.Printf("Writing blogs info to the directory '%s'", outputDir)

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	var zeroPaddings int = len(strconv.Itoa(len(blogs)))
	for i, blog := range blogs {
		blogDir := path.Join(outputDir, fmt.Sprintf("%0*d", zeroPaddings, i+1)+"-"+blog.Slug)

		log.Printf("[%d / %d] Saving '%s' as '%s'", i+1, len(blogs), blog.Title, blogDir)

		err := os.MkdirAll(blogDir, os.ModePerm)
		if err != nil {
			log.Fatalf("[ERROR] Couldn't create directory for blog: %s", err)
		}

		fileName := "info.md"
		filePath := path.Join(blogDir, fileName)
		file, err := os.Create(path.Join(blogDir, fileName))
		if err != nil {
			log.Fatalf("[WARNING] Couldn't create file '%s': %s", filePath, err)
		}
		file.WriteString("---\nTitle: " + blog.Title + "\nInfo: " + blog.Info + "\nPost Link: " + blog.PostLink + "\nThumbnail Link: " + blog.ThumbnailLink + "\nSlug: " + blog.Slug + "\n---")
		file.Close()

		fileName = "thumbnail" + filepath.Ext(blog.ThumbnailLink)
		filePath = path.Join(blogDir, fileName)
		file, err = os.Create(path.Join(blogDir, fileName))
		if err != nil {
			log.Fatalf("[WARNING] Couldn't create file '%s': %s", filePath, err)
		}

		resp, err := client.Get(blog.ThumbnailLink)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Printf("[WARNING] Couldn't download '%s': %s", blog.ThumbnailLink, err)
		}

		resp.Body.Close()
		file.Close()
	}
}

func printBlogsDetail(blogs []scrapit.Blog) {
	for i, blog := range blogs {
		fmt.Printf("[%d / %d]\n", i+1, len(blogs))
		fmt.Println(blog.Title)
		fmt.Println(blog.Info)
		fmt.Println(blog.ThumbnailLink)
		fmt.Println(blog.PostLink)
		fmt.Println(blog.Slug)
		fmt.Println()
	}
	fmt.Printf("Got %d blogs!\n", len(blogs))

}

func usage(w io.Writer) {
	flag.CommandLine.SetOutput(w)
	fmt.Fprintf(w, "Usage: %s [options] output_dir\n\nWhere options are:\n", os.Args[0])
	flag.CommandLine.PrintDefaults()
	flag.CommandLine.Name()
}

func initFlags() {
	maxPostsPage = flag.Uint("p", 0, "Maximum number of posts page to scrape the blogs from")
	toSave = flag.Bool("no-save", false, "Don't save the output")
	toPrintHelp = flag.Bool("help", false, "Print this help/usage message")

	flag.Parse()
	*toSave = !*toSave

	if *toPrintHelp {
		usage(os.Stdout)
		os.Exit(0)
	}

	nonFlagArgs := flag.Args()
	if (len(nonFlagArgs) != 1) && (*toSave) {
		fmt.Fprintf(os.Stderr, "No output directory was given. Either pass the directory as an argument or pass -no-save to not save the output.\n\n")
		usage(os.Stderr)
		os.Exit(1)
	}
	if (len(nonFlagArgs) == 1) && (!*toSave) {
		fmt.Println("[NOTE] -no-save passed with an output directory as argument")
	}
	if *toSave {
		outputDir = nonFlagArgs[0]
	}
}

func main() {
	initFlags()
	if *toSave {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			log.Fatalf("[ERROR] Couldn't work with output directory '%s' for blogs: %s", outputDir, err)
		}
	}

	blogs, err := getBlogs(*maxPostsPage, LINK, BLOG_CLASS, BLOG_INFO_CLASS, BLOG_LINK_CLASS, BLOG_STYLE_CLASS, STYLE_ATTRIB)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	printBlogsDetail(blogs)

	if !*toSave {
		os.Exit(0)
	}

	saveBlogs(blogs, outputDir)
}
