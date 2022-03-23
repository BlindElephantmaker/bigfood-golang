package table

import (
	"bigfood/internal/helpers"
	"encoding/json"
)

type Comment string

var errorTableCommentIsTooLong = helpers.NewErrorBadRequest("table comment is too long")

func (c *Comment) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	comment, err := parseComment(value)
	if err != nil {
		return err
	}

	*c = comment
	return nil
}

func NewComment() Comment {
	return ""
}

func parseComment(comment string) (Comment, error) {
	if len(comment) > 32 {
		return "", errorTableCommentIsTooLong
	}

	return Comment(comment), nil
}
