package organization

type Repository interface {
	Add(org *Organization, orgUser *User) error
}
