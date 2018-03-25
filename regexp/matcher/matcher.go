package matcher

import "unicode/utf8"

func match(regexp, text string) bool {
	firstRegexpRune, width := utf8.DecodeRuneInString(regexp)
	if firstRegexpRune == '^' {
		return matchHere(regexp[width:], text)
	}
	for textIndex := range text {
		if matchHere(regexp, text[textIndex:]) {
			return true
		}
	}
	return false
}

func matchHere(regexp, text string) bool {
	if len(regexp) == 0 {
		return true
	}
	firstRegexp, firstRegexpWidth := utf8.DecodeRuneInString(regexp)
	secondRegexp, secondRegexpWidth := utf8.DecodeRuneInString(regexp[firstRegexpWidth:])
	if secondRegexp == '*' {
		return matchStar(firstRegexp, regexp[firstRegexpWidth+secondRegexpWidth:], text)
	}
	if firstRegexp == '$' && len(regexp[firstRegexpWidth:]) == 0 {
		return len(text) == 0
	}
	firstText, firstTextWidth := utf8.DecodeRuneInString(text)
	if len(text) != 0 && (firstRegexp == '.' || firstRegexp == firstText) {
		return matchHere(regexp[firstRegexpWidth:], text[firstTextWidth:])
	}
	return false
}

func matchStar(c rune, regexp, text string) bool {
	for textIndex, textRune := range text {
		if matchHere(regexp, text[textIndex:]) {
			return true
		}
		if !(textRune == c || c == '.') {
			break
		}
	}
	return false
}
