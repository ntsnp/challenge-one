package main

import (
	"fmt"
	"log"
)

// Constants
// ---------

const LINK = "https://blog.sentry.io"
const BLOG_DIV_CLASS = ".css-spjg5j.e19gd7e57"
const BLOG_BG_DIV_CLASS = ".e19gd7e53"

// Functions
// ---------

func main() {
	blogs, err := getBlogs(LINK, BLOG_DIV_CLASS, BLOG_BG_DIV_CLASS)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.title)
		fmt.Println(blog.thumbnailLink)
	}
}
