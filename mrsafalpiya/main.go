package main

import (
	"fmt"
	"log"

	"github.com/challenge-one/scrapit"
)

// Constants
// ---------

const LINK = "https://blog.sentry.io"
const BLOG_CLASS = ".css-spjg5j.e19gd7e57"
const BLOG_LINK_CLASS = ".css-1qh9hqn.e19gd7e56"
const BLOG_STYLE_CLASS = ".e19gd7e53"
const STYLE_ATTRIB = "data-emotion"

// Functions
// ---------

func getBlogs(link string, blogClass string, blogLinkClass string, blogStyleClass string, styleAttrib string) ([]scrapit.Blog, error) {
	scrapitInstance, err := scrapit.NewScrapit(link)
	if err != nil {
		return scrapitInstance.Blogs, err
	}
	_ = scrapitInstance

	scrapitInstance.InitBlogsScrape(blogClass, blogLinkClass, blogStyleClass, styleAttrib)

	err = scrapitInstance.Run()
	if err != nil {
		return scrapitInstance.Blogs, err
	}

	return scrapitInstance.Blogs, nil
}

func main() {
	blogs, err := getBlogs(LINK, BLOG_CLASS, BLOG_LINK_CLASS, BLOG_STYLE_CLASS, STYLE_ATTRIB)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.Title)
		fmt.Println(blog.ThumbnailLink)
		fmt.Println(blog.PostLink)
	}
}
