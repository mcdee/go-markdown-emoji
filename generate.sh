#!/bin/sh

cat >emoji.go <<EOHEADER
package emoji
  
func isValidEmoji(input []byte) bool {
    _, exists := emojiMap[string(input)]
    return exists
}

var emojiMap = map[string]string{
EOHEADER

curl -s -o - https://raw.githubusercontent.com/muan/emojilib/master/emojis.json | jq -r 'to_entries | sort_by(.key) | .[] | "\"\(.key)\": \"\(.value.char)\","' >>emoji.go

cat >>emoji.go <<EOFOOTER
}
EOFOOTER

# Strip out items emojis for which we do not have graphics
for shortcode in couplekiss_woman_woman couple_with_heart_man_man female_detective weight_lifting_woman basketball_woman rainbow_flag couple_with_heart_woman_woman golfing_woman couplekiss_man_man 
do
  sed -i "/\"${shortcode}\"/d" emoji.go
done

gofmt -w emoji.go
