package emoji

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func generateURL(emoji string) string {
	code, exists := emojiMap[emoji]
	if !exists {
		return ""
	}
	res := ""
	chars := utf8.RuneCountInString(code)
	curChar := 1
	for _, c := range code {
		tmp := strings.Trim(strings.ToLower(strconv.QuoteRuneToASCII(c)), "'")
		if tmp != `\ufe0f` || (chars > 2 && curChar == chars) {
			// Valid character to add
			if curChar > 1 {
				res += "-"
			}
			if len(tmp) == 1 {
				res += fmt.Sprintf("%x", []byte(tmp))
			} else {
				res += strings.TrimLeft(tmp[2:], "0")
			}
		}
		curChar++
	}
	return fmt.Sprintf("https://twemoji.maxcdn.com/2/svg/%s.svg", res)
}
