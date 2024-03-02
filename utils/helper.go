package utils

import (
	"fmt"
	"strings"
)

func FindPath(path string) string {
	// Split the string by "/"
	parts := strings.Split(path, "/")
	fmt.Println(parts)

	return parts[2]
}
