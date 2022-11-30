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

## Output

The output creates the following directory structure:

```console
$ tree output/ | head
output/
├── 001-profiling-for-python
│   ├── info.md
│   └── thumbnail.jpg
├── 002-profiling-for-node-js
│   ├── info.md
│   └── thumbnail.jpg
├── 003-how-we-run-our-python-tests-in-hundreds-of-environments-really-fast
│   ├── info.md
│   └── thumbnail.jpg
```

Where the `info.md` consists of following data:

```md
---
Title: Application Profiling for Python
Info: Profiling is an important tool in every developer’s toolkit because it provides a granular view into the execution of your program from your…
Post Link: https://blog.sentry.io/2022/11/21/profiling-for-python/
Thumbnail Link: https://images.ctfassets.net/em6l9zw4tzag/3tMAGwsW5Z9bYPCxREi9Ay/22ed1db72eb9c241733ab20a119f5ee5/profiling-page-meta.jpg
Slug: profiling-for-python
---
```

## Frontend

A simple frontend html file is written to the output directory as `index.html`.

![Frontend](./docs/frontend.png)

## License

MIT
