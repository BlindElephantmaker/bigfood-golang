package table

import "errors"

type Comment string

var ErrorTableCommentIsTooLong = errors.New("table comment is too long")

func ParseComment(comment string) (Comment, error) {
	if len(comment) > 32 {
		return "", ErrorTableCommentIsTooLong
	}

	return Comment(comment), nil
}
