package permissions

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetPermissions(user.Id) (*Permissions, error)
}

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

const (
	cafeUserTable     = cafeUser.Table
	cafeUserRoleTable = cafeUser.RoleTable
)

type permissionValue struct {
	CafeId helpers.Uuid   `db:"cafe_id"`
	Role   sql.NullString `db:"role"`
}

func (r *RepositoryPSQL) GetPermissions(userId user.Id) (*Permissions, error) {
	query := fmt.Sprintf(`
SELECT cafe_id, role
FROM %s cu
    LEFT JOIN %s cur on cu.id = cur.cafe_user_id
WHERE user_id = $1
`, cafeUserTable, cafeUserRoleTable)

	var permissionValues []permissionValue
	err := r.db.Select(&permissionValues, query, userId)
	if err != nil {
		return nil, err
	}

	return castToPermissions(userId, &permissionValues)
}

func castToPermissions(userId user.Id, values *[]permissionValue) (*Permissions, error) {
	permissions := CreateEmptyPermission(userId)

	for _, value := range *values {
		permissions.CreateCafePerm(value.CafeId)

		if value.Role.Valid {
			permissions.AppendRole(value.CafeId, cafeUser.Role(value.Role.String))
		}
	}

	return permissions, nil
}
