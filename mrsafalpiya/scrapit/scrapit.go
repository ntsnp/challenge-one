package scrapit

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

type Blog struct {
	Title         string
	ThumbnailLink string
}

type scrapit struct {
	/* inputs */
	Link           string
	blogClass      string
	blogStyleClass string
	styleAttrib    string

	/* local states */
	collector   *colly.Collector
	backgrounds map[string]string
	protocol    string
	host        string

	/* output */
	Blogs []Blog
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
		Link:      link,
		collector: colly.NewCollector(),
		protocol:  protocol,
		host:      host,
	}, nil
}

func (s *scrapit) InitBlogsScrape(blogClass string, blogStyleClass string, styleAttrib string) {
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

		s.Blogs = append(s.Blogs, Blog{
			Title:         title,
			ThumbnailLink: bgUrl,
		})
	})
}

func (s *scrapit) Run() error {
	err := s.collector.Visit(s.Link)
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
