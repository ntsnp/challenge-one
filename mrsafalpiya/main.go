package main

import (
	"fmt"
	"log"
)

// Constants
// ---------

const LINK = "https://blog.sentry.io"
const BLOG_CLASS = ".css-spjg5j.e19gd7e57"
const BLOG_STYLE_CLASS = ".e19gd7e53"
const STYLE_ATTRIB = "data-emotion"

// Functions
// ---------

func main() {
	blogs, err := getBlogs(LINK, BLOG_CLASS, BLOG_STYLE_CLASS, STYLE_ATTRIB)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.title)
		fmt.Println(blog.thumbnailLink)
	}
}
