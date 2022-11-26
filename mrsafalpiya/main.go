package main

import "fmt"

type Blog struct {
	Title         string
	ThumbnailLink string
}

func main() {
	var blogs []Blog

	blogs = append(blogs, Blog{
		Title:         "Foo",
		ThumbnailLink: "Bar",
	})
	blogs = append(blogs, Blog{
		Title:         "Bar",
		ThumbnailLink: "Buzz",
	})

	for _, blog := range blogs {
		fmt.Println(blog.Title, "---", blog.ThumbnailLink)
	}
}
