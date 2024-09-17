package word

/*
"`~!@#$%^&*()_+=-[]{}\\|/№;:? '\"".
*/
var separators = map[rune]bool{
	'`':  true,
	'~':  true,
	'!':  true,
	'@':  true,
	'"':  true,
	'#':  true,
	'№':  true,
	';':  true,
	'$':  true,
	'%':  true,
	':':  true,
	'^':  true,
	'&':  true,
	'?':  true,
	'*':  true,
	'(':  true,
	')':  true,
	'-':  true,
	'_':  true,
	'+':  true,
	'=':  true,
	'{':  true,
	'}':  true,
	'[':  true,
	']':  true,
	'\'': true,
	'\\': true,
	'|':  true,
	'/':  true,
	' ':  true,
	'>':  true,
	'<':  true,
	'.':  true,
	',':  true,
}

func Words(s string) (wordsMap map[string]int) {
	wordsMap = make(map[string]int)
	var lastIndex int
	isword := false
	for index, symbol := range s {
		if separators[symbol] {
			if isword && lastIndex != index {
				wordsMap[s[lastIndex:index]]++
			}
			isword = false
		} else {
			if !isword {
				lastIndex = index
			}
			isword = true
		}
	}
	if isword && lastIndex != len(s) {
		wordsMap[s[lastIndex:]]++
	}
	return
}
