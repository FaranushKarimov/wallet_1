package wallet

import (
	"errors"

	"github.com/FaranushKarimov/wallet/pkg/types"
)

//ErrAccountNotFound -- account not found
var (
	ErrAccountNotFound = errors.New("account not found")
)

//Service struct
type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

// RegisterAccount function receiver
func (service *Service) RegisterAccount(phone types.Phone) {
	for _, account := range service.accounts {
		if account.Phone == phone {
			return
		}
	}
	service.nextAccountID++
	service.accounts = append(service.accounts, &types.Account{
		ID:      service.nextAccountID,
		Phone:   phone,
		Balance: 0,
	})
}

//Deposit function receiver service
func (service *Service) Deposit(accountID int64, amount types.Money) {
	if amount <= 0 {
		return
	}
	var account *types.Account

	if account == nil {
		return
	}
	account.Balance += amount
}

//FindAccountByID function
func (service *Service) FindAccountByID(accountID int64) (*types.Account, error) {

	for _, account := range service.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}
