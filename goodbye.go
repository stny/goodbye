package goodbye

import "regexp"

var (
	ansiColorCodesRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)
	blankLinesRegex      = regexp.MustCompile("(?m)^\\s*$[\r\n]*")
)

func BlankLines(s string) string {
	return blankLinesRegex.ReplaceAllString(s, "")
}

func ANSIColorCodes(s string) string {
	return ansiColorCodesRegexp.ReplaceAllString(s, "")
}
