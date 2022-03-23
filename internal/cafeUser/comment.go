package cafeUser

import (
	"encoding/json"
	"errors"
)

type Comment string

var errorCafeUserCommentIsTooLong = errors.New("cafe user comment is too long")

func NewComment() Comment {
	return ""
}

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

func parseComment(comment string) (Comment, error) {
	if len(comment) > 32 {
		return "", errorCafeUserCommentIsTooLong
	}

	return Comment(comment), nil
}
