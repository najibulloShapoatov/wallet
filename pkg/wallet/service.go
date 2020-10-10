package wallet

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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

//ErrFavoriteNotFound -- favorite not found
var ErrFavoriteNotFound = errors.New("favorite not found")

//Service model
type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
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

//Deposit method
func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount < 0 {
		return ErrAmountMustBePositive
	}
	account, err := s.FindAccountByID(accountID)
	if err != nil {
		return err
	}
	account.Balance += amount
	return nil

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

//Repeat method
func (s *Service) Repeat(paymentID string) (*types.Payment, error) {

	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}
	paymentNew, err := s.Pay(payment.AccountID, payment.Amount, payment.Category)
	if err != nil {
		return nil, err
	}
	return paymentNew, nil
}

//FavoritePayment method
func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {

	payment, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	favoriteID := uuid.New().String()
	favorite := &types.Favorite{
		ID:        favoriteID,
		AccountID: payment.AccountID,
		Name:      name,
		Amount:    payment.Amount,
		Category:  payment.Category,
	}

	s.favorites = append(s.favorites, favorite)

	return favorite, nil
}

//PayFromFavorite method
func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {

	var favorite *types.Favorite
	for _, v := range s.favorites {
		if v.ID == favoriteID {
			favorite = v
			break
		}
	}
	if favorite == nil {
		return nil, ErrFavoriteNotFound
	}

	payment, err := s.Pay(favorite.AccountID, favorite.Amount, favorite.Category)

	if err != nil {
		return nil, err
	}
	return payment, nil
}

//ExportToFile func
func (s *Service) ExportToFile(path string) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	var str string
	for _, v := range s.accounts {
		str += fmt.Sprint(v.ID) + ";" + string(v.Phone) + ";" + fmt.Sprint(v.Balance) + "|"
	}
	_, err = file.WriteString(str)

	if err != nil {
		return err
	}

	return nil
}

//ImportFromFile method
func (s *Service) ImportFromFile(path string) error {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	strArray := strings.Split(string(content), "|")
	if len(strArray) > 0 {
		strArray = strArray[:len(strArray)-1]
	}
	for _, v := range strArray {
		strArrAcount := strings.Split(v, ";")
		fmt.Println(strArrAcount)

		id, err := strconv.ParseInt(strArrAcount[0], 10, 64)
		if err != nil {
			return err
		}
		balance, err := strconv.ParseInt(strArrAcount[2], 10, 64)
		if err != nil {
			return err
		}
		account := &types.Account{
			ID:      id,
			Phone:   types.Phone(strArrAcount[1]),
			Balance: types.Money(balance),
		}
		s.accounts = append(s.accounts, account)
	}

	return nil
}
