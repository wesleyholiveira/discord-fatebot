package sanitizer

import "regexp"

const (
	RedirectRegexp       = `#Redirect (.*)`
	TitleRegexp          = ` \(.*\)`
	DoubleCurlyRegexp    = `\{\{(.*?)\}\}`
	DoubleBracketsRegexp = `\[\[(.*?)\]\]`
)

func Title(text string) string {
	re := regexp.MustCompile(TitleRegexp)
	return re.ReplaceAllString(text, "")
}

func RemoveDoubleBrackets(text string) string {
	re := regexp.MustCompile(DoubleBracketsRegexp)
	return re.ReplaceAllString(text, "$1")
}

func RemoveDoubleCurlyBraces(text string) string {
	re := regexp.MustCompile(DoubleCurlyRegexp)
	return re.ReplaceAllString(text, "$1")
}

func RemoveAll(text string) string {
	return RemoveDoubleCurlyBraces(RemoveDoubleBrackets(text))
}
