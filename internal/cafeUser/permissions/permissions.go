package permissions

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
)

type CafePermissions struct {
	Roles cafeUser.Roles
}

func createEmptyCafePermission() *CafePermissions {
	return &CafePermissions{
		Roles: cafeUser.Roles{},
	}
}

func (cp *CafePermissions) appendRole(role cafeUser.Role) {
	cp.Roles = append(cp.Roles, role)
}

type Permissions struct {
	UserId helpers.Uuid
	Cafes  map[helpers.Uuid]*CafePermissions
}

func (p *Permissions) AppendRole(cafeId helpers.Uuid, role cafeUser.Role) {
	cafePerm := p.Cafes[cafeId]
	cafePerm.appendRole(role)
}

func (p *Permissions) CreateCafePerm(cafeId helpers.Uuid) {
	_, ok := p.Cafes[cafeId]
	if !ok {
		p.Cafes[cafeId] = createEmptyCafePermission()
	}
}

func (p *Permissions) HasRole(cafeId helpers.Uuid, role cafeUser.Role) bool {
	cafePerm, ok := p.Cafes[cafeId]
	if !ok {
		return false
	}

	for _, r := range cafePerm.Roles {
		if role == r {
			return true
		}
	}
	return false
}

func CreateEmptyPermission(userId helpers.Uuid) *Permissions {
	return &Permissions{
		UserId: userId,
		Cafes:  make(map[helpers.Uuid]*CafePermissions),
	}
}
