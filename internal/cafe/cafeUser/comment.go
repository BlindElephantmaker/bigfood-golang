package cafeUser

type Comment string

// todo: is used or not
//var ErrorCafeUserCommentIsTooLong = errors.New("cafe user comment is too long")

func newComment() Comment {
	return ""
}

// todo: is used or not
//func ParseComment(comment string) (Comment, error) {
//	if len(comment) > 32 {
//		return "", ErrorCafeUserCommentIsTooLong
//	}
//
//	return Comment(comment), nil
//}
