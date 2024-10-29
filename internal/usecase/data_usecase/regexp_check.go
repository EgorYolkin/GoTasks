package data_usecase

import (
	"regexp"
)

var linkPattern = `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

func IsLink(text string) bool {

	matched, _ := regexp.MatchString(linkPattern, text)
	return matched
}
