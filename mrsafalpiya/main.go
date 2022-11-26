package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
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

	var styleBody string
	var stylesheets []*css.Stylesheet

	c := colly.NewCollector()

	c.OnHTML("style[data-emotion]", func(e *colly.HTMLElement) {
		styleBody = e.Text

		stylesheet, err := parser.Parse(styleBody)
		if err != nil {
			log.Fatalln(err)
		}
		stylesheets = append(stylesheets, stylesheet)
	})

	c.OnHTML(divClass, func(e *colly.HTMLElement) {
		title := e.ChildText("h2")

		childDivs := e.ChildAttrs("div", "class")
		bgDiv := "." + strings.Split(childDivs[len(childDivs)-1], " ")[0]
		_ = bgDiv

		var bgUrl string
		for _, style := range stylesheets {
			for _, rule := range style.Rules {
				if len(rule.Selectors) > 0 && rule.Selectors[0] == bgDiv {
					for _, decl := range rule.Declarations {
						if decl.Property == "background-image" {
							bgUrl = urlFromCSSVal(decl.Value)
						}
					}
				}
			}
		}

		blogs = append(blogs, blog{
			title:         title,
			thumbnailLink: bgUrl,
		})
	})

	c.Visit(link)

	return blogs, nil
}

/* Helper */

/*
 * Extracts the URL parameter from the CSS value.
 *
 * Example: `url(//foo)` -> `foo`
 */
func urlFromCSSVal(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			k := 0
			for ; k < len(s)-i; k++ {
				if s[i+k+1] != '/' {
					break
				}
			}
			return s[i+k+1 : j] /* the link has leading `//` */
		}
	}
	return ""
}

/* Entry point */

func main() {
	blogs, err := getBlogs(LINK, BLOG_DIV_CLASS, BLOG_BG_DIV_CLASS)
	if err != nil {
		log.Fatalf("[ERROR] Couldn't get blogs: %s", err.Error())
	}

	for _, blog := range blogs {
		fmt.Println(blog.title, blog.thumbnailLink)
	}
}
