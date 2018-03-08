package goodbye

import (
	"testing"
)

var colorcodesTestCase = [][2]string{
	{"\033[31;1;4mHello\033[0m", "Hello"},
	{"\x1b[38;2;255;100;0mTRUECOLOR\x1b[0m", "TRUECOLOR"},
}

func TestEscapeSequence(t *testing.T) {
	for _, p := range colorcodesTestCase {
		actual := ANSIColorCodes(p[0])
		expected := p[1]
		if expected != actual {
			t.Errorf("EscapeSequence(%s) = %s, want %s", p[0], actual, expected)
		}
	}
}

var blanklinesTestCase = [][3]string{
	{"A\n\nB", "A\nB"},
	{"A\n\n\nB\n\nC", "A\nB\nC"},
}

func TestBlankLines(t *testing.T) {
	for _, p := range blanklinesTestCase {
		actual := BlankLines(p[0])
		expected := p[1]
		if expected != actual {
			t.Errorf("BlankLines(...) = %s, want %s", actual, expected)
		}
	}
}
