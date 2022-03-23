package table

import (
	"bigfood/internal/helpers"
	"encoding/json"
)

type Title string

var errorTableTitleIsTooShort = helpers.NewErrorBadRequest("table title is too short")
var errorTableTitleIsTooLong = helpers.NewErrorBadRequest("table title is too long")

func (t *Title) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	title, err := ParseTitle(value)
	if err != nil {
		return err
	}

	*t = title
	return nil
}

func ParseTitle(title string) (Title, error) {
	if len(title) < 1 {
		return "", errorTableTitleIsTooShort
	}
	if len(title) > 32 {
		return "", errorTableTitleIsTooLong
	}

	return Title(title), nil
}
