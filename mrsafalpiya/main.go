package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/gocolly/colly"
)

// Constants
// ---------

const LINK = "https://blog.sentry.io"
const BLOG_DIV_CLASS = ".css-spjg5j.e19gd7e57"
const BLOG_BG_DIV_CLASS = ".e19gd7e53"

// Structs
// -------

type blog struct {
	title         string
	thumbnailLink string
}

// Functions
// ---------

/* Blogs */

func getBlogs(link string, divClass string, bgDivClass string) ([]blog, error) {
	var blogs []blog

	if !govalidator.IsURL(link) {
		return blogs, errors.New("Given link is not an URL")
	}

	c := colly.NewCollector()
	c.OnHTML(divClass, func(e *colly.HTMLElement) {
		title := e.ChildText("h2")

		blogs = append(blogs, blog{
			title:         title,
			thumbnailLink: "",
		})
	})
	c.Visit(link)

	return blogs, nil
}

/* File Handling */

/* Entry point */

func main() {
	blogs, err := getBlogs(LINK, BLOG_DIV_CLASS, BLOG_BG_DIV_CLASS)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.title)
	}
}
