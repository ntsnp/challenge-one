package main

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
	"github.com/gocolly/colly"
)

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

	u, err := url.ParseRequestURI(link)
	if err != nil {
		return blogs, err
	}
	protocol := u.Scheme + "://"
	host := u.Host

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
							bgUrl = cleanUrl(bgUrl, protocol, host)
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

	err = c.Visit(link)
	if err != nil {
		return blogs, err
	}

	return blogs, nil
}
