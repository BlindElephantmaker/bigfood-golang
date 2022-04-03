package createCafe

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"bigfood/pkg/database"
)

type Response struct {
	Cafe *cafe.Cafe `json:"cafe"`
}

func (h *Handler) Run(userId user.Id) (*Response, error) {
	newCafe := cafe.New()
	newCafeUser := cafeUser.NewCafeUser(newCafe.Id, userId, cafeUser.NewComment())
	now := helpers.NowTime()

	tx, err := h.Transactions.Begin()
	if err != nil {
		return nil, err
	}

	if err := h.cafeRepository.AddTx(tx, newCafe, now); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := h.cafeUserRepository.AddTx(tx, newCafeUser, now, cafeUser.Roles{
		cafeUser.Owner,
		cafeUser.Admin,
		cafeUser.Hostess,
	}); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return &Response{newCafe}, nil
}

type Handler struct {
	cafeRepository     cafe.Repository
	cafeUserRepository cafeUser.Repository
	Transactions       *database.TransactionFactory
}

func New(
	cafes cafe.Repository,
	cafeUsers cafeUser.Repository,
	transactions *database.TransactionFactory,
) *Handler {
	return &Handler{
		cafeRepository:     cafes,
		cafeUserRepository: cafeUsers,
		Transactions:       transactions,
	}
}
