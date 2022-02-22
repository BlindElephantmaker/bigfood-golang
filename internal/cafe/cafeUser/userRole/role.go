package userRole

type Role string

type Roles []Role

const (
	Owner   Role = "owner"
	Admin   Role = "admin"
	Hostess Role = "hostess"
)
