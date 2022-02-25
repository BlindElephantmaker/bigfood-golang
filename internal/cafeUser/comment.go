package cafeUser

import "errors"

type Comment string

var ErrorCafeUserCommentIsTooLong = errors.New("cafe user comment is too long")

func NewComment() Comment {
	return ""
}

func ParseComment(comment string) (Comment, error) {
	if len(comment) > 32 {
		return "", ErrorCafeUserCommentIsTooLong
	}

	return Comment(comment), nil
}
