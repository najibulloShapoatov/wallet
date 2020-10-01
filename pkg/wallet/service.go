package wallet

import (
	"errors"

	"github.com/google/uuid"
	"github.com/najibulloShapoatov/wallet/pkg/types"
)

//ErrPhoneRegistered -- phone already registred
var ErrPhoneRegistered = errors.New("phone already registred")

//ErrAmountMustBePositive -- amount must be greater than zero
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")

//ErrAccountNotFound -- account not found
var ErrAccountNotFound = errors.New("account not found")

//ErrNotEnoughtBalance -- account not found
var ErrNotEnoughtBalance = errors.New("account not enought balance")

//ErrPaymentNotFound -- account not found
var ErrPaymentNotFound = errors.New("payment not found")

//Service model
type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

//RegisterAccount meth
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)

	return account, nil
}

//Pay method
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {

	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}
	var account *types.Account
	for _, ac := range s.accounts {
		if ac.ID == accountID {
			account = ac
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughtBalance
	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

//FindAccountByID method
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {

	for _, account := range s.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}

//FindPaymentByID method
func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {

	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}

//Reject method
func (s *Service) Reject(paymentID string) error {

	var payment, err = s.FindPaymentByID(paymentID)

	if err != nil {
		return err
	}

	var account, er = s.FindAccountByID(payment.AccountID)

	if er != nil {
		return er
	}

	payment.Status = types.PaymentStatusFail
	account.Balance += payment.Amount
	
	return nil
}
