package foodchain

import (
	"fmt"
	"strings"
)

// Verse will print the single verse requested
func Verse(n int) string {
	upTo := n - 1
	ret := []string{fmt.Sprintf("I know an old lady who swallowed a %s.", animals[upTo].name)}

	if animals[upTo].action != "" {
		ret = append(ret, animals[upTo].action)
	}

	if upTo != len(animals)-1 {
		for i := upTo; i > 0; i-- {
			ret = append(ret, animals[i].resolve)
		}
		ret = append(ret, die)
	} else {
		ret = append(ret, dead)
	}

	return strings.Join(ret, "\n")
}

// Verses will print the verses from start to end
func Verses(start, end int) string {
	song := []string{}
	for i := start; i < end; i++ {
		song = append(song, Verse(i))
	}
	song = append(song, Verse(end))
	return strings.Join(song, "\n\n")
}

// Song will print out the entire song
func Song() string {
	return Verses(1, len(animals))
}
