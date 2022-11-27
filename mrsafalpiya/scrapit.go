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

type scrapit struct {
	/* inputs */
	link           string
	blogClass      string
	blogStyleClass string
	styleAttrib    string

	/* local states */
	collector   *colly.Collector
	backgrounds map[string]string
	protocol    string
	host        string

	/* output */
	blogs []blog
}

// Methods
// -------

/* scrapitSentry */

func NewScrapit(link string) (*scrapit, error) {
	u, err := url.ParseRequestURI(link)
	if err != nil {
		return nil, err
	}
	protocol := u.Scheme + "://"
	host := u.Host

	if !govalidator.IsURL(link) {
		return nil, errors.New("Given link is not an URL")
	}

	return &scrapit{
		link:      link,
		collector: colly.NewCollector(),
		protocol:  protocol,
		host:      host,
	}, nil
}

func (s *scrapit) initBlogsScrape(blogClass string, blogStyleClass string, styleAttrib string) {
	s.blogClass = blogClass
	s.blogStyleClass = blogStyleClass
	s.styleAttrib = styleAttrib
	s.backgrounds = make(map[string]string)

	/* style attrib handler */

	s.collector.OnHTML("style["+s.styleAttrib+"]", func(e *colly.HTMLElement) {
		styleBody := e.Text

		stylesheet, err := parser.Parse(styleBody)
		if err != nil {
			log.Fatalln(err)
		}
		s.addBgs(stylesheet)
	})

	/* blog div handler */

	s.collector.OnHTML(s.blogClass, func(e *colly.HTMLElement) {
		title := e.ChildText("h1, h2, h3")

		childDivs := e.ChildAttrs("div", "class")
		bgDiv := "." + strings.Split(childDivs[len(childDivs)-1], " ")[0]
		bgUrl := cleanUrl(urlFromCSSVal(s.backgrounds[bgDiv]), s.protocol, s.host)

		s.blogs = append(s.blogs, blog{
			title:         title,
			thumbnailLink: bgUrl,
		})
	})
}

func (s *scrapit) run() error {
	err := s.collector.Visit(s.link)
	if err != nil {
		return err
	}

	return nil
}

func (s *scrapit) addBgs(stylesheet *css.Stylesheet) {
	for _, rule := range stylesheet.Rules {
		for _, decl := range rule.Declarations {
			if len(rule.Selectors) > 0 && decl.Property == "background-image" {
				s.backgrounds[rule.Selectors[0]] = decl.Value
			}
		}
	}
}

// Functions
// ---------

func getBlogs(link string, blogClass string, blogStyleClass string, styleAttrib string) ([]blog, error) {
	scrapitInstance, err := NewScrapit(link)
	if err != nil {
		return scrapitInstance.blogs, err
	}
	_ = scrapitInstance

	scrapitInstance.initBlogsScrape(blogClass, blogStyleClass, styleAttrib)

	err = scrapitInstance.run()
	if err != nil {
		return scrapitInstance.blogs, err
	}

	return scrapitInstance.blogs, nil
}
