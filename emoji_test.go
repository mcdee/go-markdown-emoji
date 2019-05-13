package emoji

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
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
		{ // 0
			input:  []byte(``),
			output: []byte(``),
		},
		{ // 1
			input:  []byte(`:`),
			output: []byte(`<p>:</p>`),
		},
		{ // 2
			input:  []byte(`::`),
			output: []byte(`<p>::</p>`),
		},
		{ // 3
			input:  []byte(`notemoji`),
			output: []byte(`<p>notemoji</p>`),
		},
		{ // 4
			input:  []byte(`:smile:`),
			output: []byte(`<p><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img></p>`),
		},
		{ // 5
			input:  []byte(`:smile: :smile:`),
			output: []byte(`<p><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img> <img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img></p>`),
		},
		{ // 6
			input:  []byte(`:smile::smile:`),
			output: []byte(`<p><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img></p>`),
		},
		{ // 7
			input:  []byte(`A big :smile:`),
			output: []byte(`<p>A big <img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img></p>`),
		},
		{ // 8
			input:  []byte(`:smile: for you`),
			output: []byte(`<p><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img> for you</p>`),
		},
		{ // 9
			input:  []byte(`A *big* :smile: for you`),
			output: []byte(`<p>A <em>big</em> <img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f604.svg" alt=":smile:"></img> for you</p>`),
		},
		{ // 10
			input:  []byte(`:woman_cartwheeling:`),
			output: []byte(`<p><img class="emoji" src="https://twemoji.maxcdn.com/2/svg/1f938-200d-2640-fe0f.svg" alt=":woman_cartwheeling:"></img></p>`),
		},
		{ // 11
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

func TestValidLinks(t *testing.T) {
	for k := range emojiMap {
		url := generateURL(k)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to obtain %s (%s): %v\n", k, url, err)
		}
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Failed to obtain %s (%s): %v\n", k, url, err)
		}
	}
}

func TestGenerateReference(t *testing.T) {
	cols := 4

	mapKeys := make([]string, 0)
	for k := range emojiMap {
		mapKeys = append(mapKeys, k)
	}
	sort.Strings(mapKeys)
	i := 0
	for _, k := range mapKeys {
		if i%cols == 0 {
			fmt.Printf("<tr>")
		}
		fmt.Printf("<td>:%s:</td><td><img src=\"%s\"></img></td>", k, generateURL(k))
		if i%cols == cols-1 {
			fmt.Printf("</tr>\n")
		}
		i++
	}
	if i%cols != 0 {
		fmt.Printf("</tr>\n")
	}
}
