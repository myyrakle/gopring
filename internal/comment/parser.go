package comment

import "strings"

// FIXME: 엣지케이스 고려 안하고 대충짬
func ParseCommentBlocks(text string) []string {
	comments := []string{}

	inBlockComment := false
	blockCommentBuffer := ""
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)

		if inBlockComment {
			if strings.HasSuffix(line, "*/") {
				inBlockComment = false
				blockCommentBuffer += strings.TrimSpace(line[:len(line)-2])
				comments = append(comments, blockCommentBuffer)
				continue
			}

			blockCommentBuffer += line
			continue
		}

		if strings.HasPrefix(line, "/*") {
			// 한 줄에서 끝나는 경우
			if strings.HasSuffix(line, "*/") {
				comments = append(comments, strings.TrimSpace(line[2:len(line)-2]))
				continue
			}

			// 주석이 여러 줄에 걸쳐 있는 경우
			inBlockComment = true
			blockCommentBuffer = strings.TrimSpace(line[2:])
			continue
		}

		if strings.HasPrefix(line, "//") {
			comments = append(comments, strings.TrimSpace(line[2:]))
			continue
		}
	}

	return comments
}
