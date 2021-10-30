package internal

import "strings"

func GetDocId(username string) string {
	tokens := strings.Split(username, "@")
	return tokens[0]
}
