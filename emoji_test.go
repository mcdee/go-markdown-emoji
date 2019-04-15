package emoji

import (
	"bytes"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func TestEmoji(t *testing.T) {
	tests := []struct {
		input  []byte
		output []byte
	}{
		{
			input:  []byte(``),
			output: []byte(``),
		},
		{
			input:  []byte(`:`),
			output: []byte(`<p>:</p>`),
		},
		{
			input:  []byte(`::`),
			output: []byte(`<p>::</p>`),
		},
		{
			input:  []byte(`notemoji`),
			output: []byte(`<p>notemoji</p>`),
		},
		{
			input:  []byte(`:smile:`),
			output: []byte(`<img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img>`),
		},
		{
			input:  []byte(`:basketball_woman:`),
			output: []byte(`<img class="emoji" src="https://twemoji.maxcdn.com/2/svg/26f9-fe0f-200d-2640-fe0f.svg" alt=":basketball_woman:"></img>`),
		},
		{
			input:  []byte(`:unknown:`),
			output: []byte(`<p>:unknown:</p>`),
		},
	}

	for i, test := range tests {
		p := parser.New()
		p.Opts = parser.Options{ParserHook: Parser}
		r := html.NewRenderer(html.RendererOptions{RenderNodeHook: Renderer})
		output := bytes.TrimSpace(markdown.ToHTML(test.input, p, r))
		if bytes.Compare(output, test.output) != 0 {
			t.Errorf("failed at test %d\nExpected: %s\nReceived: %s\n", i, test.output, output)
		}
	}
}
