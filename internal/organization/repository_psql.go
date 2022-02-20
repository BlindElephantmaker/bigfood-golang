package organization

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryPSQL struct {
	db *sqlx.DB
}

const (
	organizationTable         = "organization"
	organizationUserTable     = "organization_user"
	organizationUserRoleTable = "organization_user_role"
)

func NewRepositoryPSQL(db *sqlx.DB) *RepositoryPSQL {
	return &RepositoryPSQL{db: db}
}

func (repo *RepositoryPSQL) Add(org *Organization, orgUser *User) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	fmt.Println(org.Id, orgUser.UserId)

	queryOrg := fmt.Sprintf("INSERT INTO %s (id, created_at) VALUES ($1, $2)", organizationTable)
	result, err := tx.Exec(queryOrg, org.Id.String(), org.CreatedAt)
	fmt.Println(result)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	queryOrgUser := fmt.Sprintf("INSERT INTO %s (id, organization_id, user_id, created_at) VALUES ($1, $2, $3, $4)",
		organizationUserTable)
	_, err = tx.Exec(queryOrgUser, orgUser.Id.String(), orgUser.OrganizationId.String(), orgUser.UserId.String(), orgUser.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	queryOrgUserRole := fmt.Sprintf("INSERT INTO %s (organization_user_id, role) VALUES ($1, $2)", organizationUserRoleTable)
	for _, role := range orgUser.Roles {
		_, err = tx.Exec(queryOrgUserRole, orgUser.Id.String(), role)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
