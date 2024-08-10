package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != HTTP_PREFIX {
		return HTTP + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, HTTP, "", 1)
	newURL = strings.Replace(newURL, HTTPS, "", 1)
	newURL = strings.Replace(newURL, WWW, "", 1)

	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}
