package cafeUser

import (
	"bigfood/internal/cafeUser/role"
	"bigfood/internal/helpers"
	"database/sql"
	"fmt"
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

func (r *RepositoryPSQL) Get(cafeId, userId helpers.Uuid) (*User, error) {
	var cafePSQL struct {
		Id        helpers.Uuid `db:"id"`
		CafeId    helpers.Uuid `db:"cafe_id"`
		UserId    helpers.Uuid `db:"user_id"`
		Comment   Comment      `db:"comment"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
	queryCafeUser := fmt.Sprintf(
		"SELECT id, cafe_id, user_id, comment, deleted_at FROM %s WHERE cafe_id = $1 AND user_id = $2",
		cafeUserTable,
	)
	err := r.db.Get(&cafePSQL, queryCafeUser, cafeId, userId)
	if err != nil {
		return nil, err
	}

	var cafeUserRoles role.Roles
	queryUserRoles := fmt.Sprintf("SELECT role FROM %s WHERE cafe_user_id = $1", cafeUserRoleTable)
	if err := r.db.Get(&cafeUserRoles, queryUserRoles, cafePSQL.UserId); err == ErrorNoResult {
		cafeUserRoles = role.Roles{}
	} else if err != nil {
		return nil, err
	}

	var deleteAt *time.Time
	if cafePSQL.DeletedAt.Valid {
		deleteAt = &cafePSQL.DeletedAt.Time
	} else {
		deleteAt = nil
	}

	return CreateCafeUser(
		cafePSQL.Id,
		cafePSQL.CafeId,
		cafePSQL.UserId,
		cafePSQL.Comment,
		deleteAt,
		cafeUserRoles,
	), nil
}

func (r *RepositoryPSQL) Add(cafeUser *User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := r.AddTx(tx, cafeUser, helpers.TimeNow()); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RepositoryPSQL) Update(cafeUser *User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := r.UpdateTx(tx, cafeUser); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RepositoryPSQL) GetUserPermissions(userId helpers.Uuid) (*role.Permissions, error) {
	query := fmt.Sprintf(`
SELECT id, cafe_id, user_id, role
FROM %s cu
    LEFT JOIN %s cur on cu.id = cur.cafe_user_id
WHERE user_id = $1
`, cafeUserTable, cafeUserRoleTable)

	var permissionValues []permissionValues
	err := r.db.Select(&permissionValues, query, userId)
	if err != nil {
		return nil, err
	}

	return castToPermissions(userId, &permissionValues)
}

type permissionValues struct {
	Id     string         `db:"id"`
	CafeId helpers.Uuid   `db:"cafe_id"`
	UserId string         `db:"user_id"`
	Role   sql.NullString `db:"role"`
}

func castToPermissions(userId helpers.Uuid, values *[]permissionValues) (*role.Permissions, error) {
	permissions := role.CreateEmptyPermission(userId)

	for _, value := range *values {
		permissions.CreateCafePerm(value.CafeId)

		if value.Role.Valid {
			permissions.AppendRole(value.CafeId, role.Role(value.Role.String))
		}
	}

	return permissions, nil
}

func (r *RepositoryPSQL) AddTx(tx *sql.Tx, cafeUser *User, createAt time.Time) error {
	queryCafeUser := fmt.Sprintf("INSERT INTO %s (id, cafe_id, user_id, comment, created_at) VALUES ($1, $2, $3, $4, $5)",
		cafeUserTable)
	_, err := tx.Exec(queryCafeUser, cafeUser.Id, cafeUser.CafeId, cafeUser.UserId, cafeUser.Comment, createAt)
	if err != nil {
		return err
	}

	return r.AddRoleTx(tx, cafeUser)
}

func (r *RepositoryPSQL) AddRoleTx(tx *sql.Tx, cafeUser *User) error {
	query := fmt.Sprintf("INSERT INTO %s (cafe_user_id, role) VALUES ($1, $2)", cafeUserRoleTable)
	for _, userRole := range cafeUser.Roles {
		_, err := tx.Exec(query, cafeUser.Id, userRole)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RepositoryPSQL) UpdateTx(tx *sql.Tx, cafeUser *User) error {
	if _, err := tx.Exec(
		fmt.Sprintf("UPDATE %s SET comment = $2, deleted_at = $3 WHERE id = $1", cafeUserTable),
		cafeUser.Id,
		cafeUser.Comment,
		cafeUser.DeletedAt,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE cafe_user_id = $1", cafeUserRoleTable),
		cafeUser.Id,
	); err != nil {
		return err
	}

	return r.AddRoleTx(tx, cafeUser)
}
