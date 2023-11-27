package util

import (
	"regexp"
)

var formatUrl = "(http|https):([a-z0-9]+).(com|info|edu)"

func IsValidUrl(url string) bool {
	r, err := regexp.Compile(formatUrl)
	if err != nil {
		return false
	}

	return r.MatchString(url)
}
