package table

import "errors"

type Title string

var ErrorTableTitleIsTooShort = errors.New("table title is too short")
var ErrorTableTitleIsTooLong = errors.New("table title is too long")

func NewTitle(title string) (Title, error) {
	if len(title) < 1 {
		return "", ErrorTableTitleIsTooShort
	}
	if len(title) > 32 {
		return "", ErrorTableTitleIsTooLong
	}

	return Title(title), nil
}
