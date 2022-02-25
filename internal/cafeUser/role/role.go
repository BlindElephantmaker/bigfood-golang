package role

import "errors"

const (
	Owner   Role = "owner"
	Admin   Role = "admin"
	Hostess Role = "hostess"
)

type Role string
type Roles []Role

var ErrorUserRoleInvalid = errors.New("user role invalid")

func parseRole(value string) (Role, error) {
	role := Role(value)
	switch role {
	case Owner, Admin, Hostess:
		return role, nil
	}

	return "", ErrorUserRoleInvalid
}

func ParseRoles(values []string) (Roles, error) {
	var roles Roles
	for _, value := range values {
		role, err := parseRole(value)
		if err != nil {
			return Roles{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
