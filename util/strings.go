package util

func IndentAfter(s, prefix string, start int) string {
	return string(IndentBytesAfter([]byte(s), []byte(prefix), start))
}

// IndentBytesAfter indents bytes after a specific line number with a given prefix.
// It returns a new byte slice with the indented contents.
func IndentBytesAfter(b, prefix []byte, start int) []byte {
	if len(b) == 0 {
		return b
	}

	bol := true
	line := 1

	res := make([]byte, 0, len(b)+len(prefix)*4)
	for _, c := range b {
		if bol && c != '\n' {
			if line > start {
				res = append(res, prefix...)
			}
		}
		res = append(res, c)
		bol = c == '\n'
		if bol {
			line++
		}
	}
	return res
}
