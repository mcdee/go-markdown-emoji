package emoji

import (
	"github.com/gomarkdown/markdown/ast"
)

// Node is a node containing an emoji
type Node struct {
	ast.Leaf
}

// Parser is a Markdown parser hook to parse emojis
func Parser(data []byte) (ast.Node, []byte, int) {
	dataLen := len(data)
	if dataLen <= 1 {
		// Not long enough to be an emoji
		return nil, nil, 0
	}
	if data[0] != ':' {
		// Incorrect start delimiter
		return nil, nil, 0
	}
	end := 1
	for ; end < dataLen; end++ {
		if data[end] == ':' {
			break
		}
	}
	if end == 1 {
		// Didn't find end delimiter or emoji is empty (and so not an emoji)
		return nil, nil, 0
	}
	if !isValidEmoji(data[1:end]) {
		// Unknown emoji
		return nil, nil, 0
	}
	node := &Node{ast.Leaf{Literal: data[1:end]}}
	return node, nil, end + 1
}
