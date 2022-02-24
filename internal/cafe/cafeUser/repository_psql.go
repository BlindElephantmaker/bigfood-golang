package cafeUser

import (
	"bigfood/internal/cafe/cafeUser/role"
	"bigfood/internal/helpers"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryPSQL struct {
	db *sqlx.DB
}

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

const (
	cafeUserTable     = "cafe_user"
	cafeUserRoleTable = "cafe_user_role"
)

type permissionValues struct {
	Id     string         `db:"id"`
	CafeId helpers.Uuid   `db:"cafe_id"`
	UserId string         `db:"user_id"`
	Role   sql.NullString `db:"role"`
}

func castToPermissions(userId *uuid.UUID, values *[]permissionValues) (*role.Permissions, error) {
	permissions := role.CreateEmptyPermission(helpers.Uuid(userId.String()))

	for _, value := range *values {
		permissions.CreateCafePerm(value.CafeId)

		if value.Role.Valid {
			permissions.AppendRole(value.CafeId, role.Role(value.Role.String))
		}
	}

	return permissions, nil
}

func (r *RepositoryPSQL) GetUserPermissions(userId *uuid.UUID) (*role.Permissions, error) {
	query := fmt.Sprintf(`
SELECT id, cafe_id, user_id, role
FROM %s cu
    LEFT JOIN %s cur on cu.id = cur.cafe_user_id
WHERE user_id = $1
`, cafeUserTable, cafeUserRoleTable)

	var permissionValues []permissionValues
	err := r.db.Select(&permissionValues, query, userId.String())
	if err != nil {
		return nil, err
	}

	return castToPermissions(userId, &permissionValues)
}

func (r *RepositoryPSQL) AddTx(tx *sql.Tx, cafeUser *User, createAt *time.Time) error {
	queryCafeUser := fmt.Sprintf("INSERT INTO %s (id, cafe_id, user_id, comment, created_at) VALUES ($1, $2, $3, $4, $5)",
		cafeUserTable)
	_, err := tx.Exec(queryCafeUser, cafeUser.Id.String(), cafeUser.CafeId.String(), cafeUser.UserId.String(), cafeUser.Comment, createAt)
	if err != nil {
		return err
	}

	queryCafeUserRole := fmt.Sprintf("INSERT INTO %s (cafe_user_id, role) VALUES ($1, $2)", cafeUserRoleTable)
	for _, userRole := range cafeUser.Roles {
		_, err = tx.Exec(queryCafeUserRole, cafeUser.Id.String(), userRole)
		if err != nil {
			return err
		}
	}

	return nil
}
