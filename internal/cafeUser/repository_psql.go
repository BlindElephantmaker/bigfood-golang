package cafeUser

import (
	"bigfood/internal/cafe"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryPsql struct {
	db *sqlx.DB
}

func NewRepositoryPsql(db *sqlx.DB) *RepositoryPsql {
	return &RepositoryPsql{db: db}
}

const (
	Table     = "cafe_user"
	RoleTable = "cafe_user_role"

	excludeRole = Owner
)

func (r *RepositoryPsql) Get(cafeUserId Id) (*CafeUser, error) {
	var cafeUser CafeUser
	query := fmt.Sprintf("SELECT id, cafe_id, user_id, comment, deleted_at FROM %s WHERE id = $1", Table)
	err := r.db.Get(&cafeUser, query, cafeUserId)
	if err != nil {
		return nil, err
	}

	return &cafeUser, nil
}

func (r *RepositoryPsql) GetByCafe(cafeId cafe.Id) ([]*CafeUser, error) {
	var cafeUsers []*CafeUser
	query := fmt.Sprintf(
		"SELECT id, cafe_id, user_id, comment, deleted_at FROM %s WHERE cafe_id = $1 AND deleted_at IS NULL",
		Table,
	)
	err := r.db.Select(&cafeUsers, query, cafeId)
	if err != nil && err != ErrorNoResult {
		return nil, err
	}

	return cafeUsers, nil
}

func (r *RepositoryPsql) GetByCafeAndUser(cafeId cafe.Id, userId user.Id) (*CafeUser, error) {
	var cafeUser CafeUser
	queryCafeUser := fmt.Sprintf(
		"SELECT id, cafe_id, user_id, comment, deleted_at FROM %s WHERE cafe_id = $1 AND user_id = $2",
		Table,
	)
	err := r.db.Get(&cafeUser, queryCafeUser, cafeId, userId)
	if err != nil {
		return nil, err
	}

	return &cafeUser, nil
}

func (r *RepositoryPsql) GetUserRoles(cafeUserId Id) (*Roles, error) {
	var roles Roles
	query := fmt.Sprintf("SELECT role FROM %s WHERE cafe_user_id = $1 and role != $2", RoleTable)
	err := r.db.Select(&roles, query, cafeUserId, excludeRole)
	if err == ErrorNoResult {
		roles = Roles{}
	} else if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (r *RepositoryPsql) Add(cafeUser *CafeUser, roles Roles) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := r.AddTx(tx, cafeUser, helpers.NowTime(), roles); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RepositoryPsql) Update(cafeUser *CafeUser, roles Roles) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err := r.UpdateTx(tx, cafeUser, roles); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RepositoryPsql) Delete(cafeUserId Id) error {
	now := helpers.NowTime()
	query := fmt.Sprintf("UPDATE %s SET deleted_at = $2 WHERE id = $1 AND deleted_at IS NULL", Table)
	_, err := r.db.Exec(query, cafeUserId, now)
	return err
}

func (r *RepositoryPsql) AddTx(tx *sql.Tx, cafeUser *CafeUser, createAt time.Time, roles Roles) error {
	query := fmt.Sprintf("INSERT INTO %s (id, cafe_id, user_id, comment, created_at) VALUES ($1, $2, $3, $4, $5)", Table)
	_, err := tx.Exec(query, cafeUser.Id, cafeUser.CafeId, cafeUser.UserId, cafeUser.Comment, createAt)
	if err != nil {
		return err
	}

	return r.AddRolesTx(tx, cafeUser, roles)
}

func (r *RepositoryPsql) AddRolesTx(tx *sql.Tx, cafeUser *CafeUser, roles Roles) error {
	query := fmt.Sprintf("INSERT INTO %s (cafe_user_id, role) VALUES ($1, $2)", RoleTable)
	for _, userRole := range roles {
		_, err := tx.Exec(query, cafeUser.Id, userRole)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RepositoryPsql) UpdateTx(tx *sql.Tx, cafeUser *CafeUser, roles Roles) error {
	if _, err := tx.Exec(
		fmt.Sprintf("UPDATE %s SET comment = $2, deleted_at = $3 WHERE id = $1", Table),
		cafeUser.Id,
		cafeUser.Comment,
		cafeUser.DeletedAt,
	); err != nil {
		return err
	}

	if _, err := tx.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE cafe_user_id = $1 AND role != $2", RoleTable),
		cafeUser.Id,
		excludeRole,
	); err != nil {
		return err
	}

	return r.AddRolesTx(tx, cafeUser, roles)
}
