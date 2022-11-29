# `get-sentry-blogs`

This is my attempt to [Challenge-one by
ntsnp](https://github.com/ntsnp/challenge-one) in [golang](https://go.dev/).

## Installation

Assuming [proper golang setup](https://go.dev/doc/install), simply clone this
repo, cd to repo's current directory and run the following command:

```console
$ go install
```

## Usage

```console
$ get-sentry-blogs -help
Usage: ./get-sentry-blogs [options] output_dir
 
Where options are:
  -help
        Print this help/usage message
  -no-save
        Don't save the output
  -p uint
        Maximum number of posts page to scrape the blogs from
```

NOTE: By default all blog posts are scraped -- this may take a while. Refer to
`-p` flag to control this.

## Frontend

A simple frontend html file is written to the output directory as `index.html`.

![Frontend](./docs/frontend.png)

## License

MIT
