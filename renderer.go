package emoji

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

// Renderer is a renderer hook to show emojis
func Renderer(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// Only handle emoji nodes
	if _, ok := node.(*Node); !ok {
		return ast.GoToNext, false
	}

	// Only handle when entering (we should only ever enter, as we're a leaf
	// node, but leave this just in case of upstream changes)
	if !entering {
		return ast.GoToNext, true
	}

	name := string(node.(*Node).Literal)
	code := emojiMap[name]
	res := ""
	for i, c := range code {
		if i > 0 {
			res += "-"
		}
		tmp := strings.ToLower(strconv.QuoteRuneToASCII(c))
		res += strings.TrimLeft(tmp[3:len(tmp)-1], "0")
	}
	url := fmt.Sprintf("https://twemoji.maxcdn.com/2/svg/%s.svg", res)
	w.Write([]byte(fmt.Sprintf(`<img class="emoji" src="%s" alt=":%s:"></img>`, url, name)))

	return ast.GoToNext, true
}
