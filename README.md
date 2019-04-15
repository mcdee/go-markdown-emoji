# go-markdown-emoji

[![Tag](https://img.shields.io/github/tag/mcdee/go-markdown-emoji.svg)](https://github.com/mcdee/go-markdown-emoji/releases/)
[![License](https://img.shields.io/github/license/mcdee/go-markdown-emoji.svg)](LICENSE)
[![Travis CI](https://img.shields.io/travis/mcdee/go-markdown-emoji.svg)](https://travis-ci.org/mcdee/go-markdown-emoji)
[![codecov.io](https://img.shields.io/codecov/c/github/mcdee/go-markdown-emoji.svg)](https://codecov.io/github/mcdee/go-markdown-emoji)

Go module to add emoji support to [go-markdown][https://github.com/gomarkdown/markdown].


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-markdown-emoji` is a standard Go module which can be installed with:

```sh
go get github.com/mcdee/go-markdown-emoji
```

## Usage

`go-markdown-emoji` provides parser and renderer hooks to the markdown engine.  The parser hook is `Parser` and the renderer `Renderer`.

Emojis are signified in Markdown as names between colons, for example `:smile:`.  A full list of the emojis supported can be seen in `emoji.go`.

### Example

```go
package main

import (
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/parser"
    emoji "github.com/mcdee/go-markdown-emoji"
)

func main() {
    p := parser.New()
    p.Opts = parser.Options{ParserHook: emoji.Parser}
    r := html.NewRenderer(html.RendererOptions{
      Flags: html.CommonFlags,
      RenderNodeHook: emoji.Renderer,
    })

    html := markdown.ToHTML(":smile:", p, r)
    fmt.Printf("%s\n", string(html))
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/mcdee/go-markdown-emoji/issues).

## License

[Apache-2.0](LICENSE).
