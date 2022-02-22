package userRole

import (
	"bigfood/internal/helpers"
)

// todo: maybe move to authorization
type CafePermissions struct {
	Roles Roles
}

func createEmptyCafePermission() *CafePermissions {
	return &CafePermissions{
		Roles: Roles{},
	}
}

func (cp *CafePermissions) appendRole(role Role) {
	cp.Roles = append(cp.Roles, role)
}

type Permissions struct {
	UserId helpers.Uuid
	Cafes  map[helpers.Uuid]*CafePermissions
}

func (p *Permissions) AppendRole(cafeId helpers.Uuid, role Role) {
	cafePerm := p.Cafes[cafeId]
	cafePerm.appendRole(role)
}

func (p *Permissions) CreateCafePerm(cafeId helpers.Uuid) {
	_, ok := p.Cafes[cafeId]
	if !ok {
		p.Cafes[cafeId] = createEmptyCafePermission()
	}
}

func CreateEmptyPermission(userId helpers.Uuid) *Permissions {
	return &Permissions{
		UserId: userId,
		Cafes:  make(map[helpers.Uuid]*CafePermissions),
	}
}
