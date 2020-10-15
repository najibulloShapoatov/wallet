package wallet

import (
	"testing"
)

func TestService_FindAccountByID_success_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method returned not nil error, account => %v", account)
	}

}

func TestService_FindAccountByID_notFound_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(2)

	if err == nil {
		t.Errorf("method returned nil error, account => %v", account)
	}

}

func TestService_Reject_success_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992000000001")
	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, error => %v", err)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "Cafe")

	if err != nil {
		t.Errorf("method Pay returned not nil error, error => %v", err)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, error => %v", err)
	}

	err = svc.Reject(pay.ID)

	if err != nil {
		t.Errorf("method Reject returned not nil error, error => %v", err)
	}

}

func TestService_Reject_fail_user(t *testing.T) {
	var svc Service
	svc.RegisterAccount("+992000000001")
	account, err := svc.FindAccountByID(1)

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "Cafe")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	err = svc.Reject(pay.ID + "uu")

	if err == nil {
		t.Errorf("method Reject returned not nil error, pay => %v", pay)
	}

}
func TestService_Repeat_success_user(t *testing.T) {
	var svc Service

	account, err := svc.RegisterAccount("+992000000001")

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "Cafe")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	paymentNew, err := svc.Repeat(pay.ID)

	if err != nil {
		t.Errorf("method Repat returned not nil error, paymentNew => %v", paymentNew)
	}

}
func TestService_Favorite_success_user(t *testing.T) {
	var svc Service

	account, err := svc.RegisterAccount("+992000000001")

	if err != nil {
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit returned not nil error, error => %v", err)
	}

	payment, err := svc.Pay(account.ID, 10_00, "Cafe")

	if err != nil {
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	favorite, err := svc.FavoritePayment(payment.ID, "My Favorite")

	if err != nil {
		t.Errorf("method FavoritePayment returned not nil error, favorite => %v", favorite)
	}

	paymentFavorite, err := svc.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Errorf("method PayFromFavorite returned not nil error, paymentFavorite => %v", paymentFavorite)
	}

}

func TestService_Export_success_user(t *testing.T) {
	var svc Service

	svc.RegisterAccount("+992000000001")
	svc.RegisterAccount("+992000000002")
	svc.RegisterAccount("+992000000003")

	err := svc.ExportToFile("export.txt")
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

}

func TestService_Import_success_user(t *testing.T) {
	var svc Service


	err := svc.ImportFromFile("export.txt")
	
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

}



func TestService_ExportImport_success_user(t *testing.T) {
	var svc Service

	svc.RegisterAccount("+992000000001")
	svc.RegisterAccount("+992000000002")
	svc.RegisterAccount("+992000000003")
	svc.RegisterAccount("+992000000004")
	
	err := svc.Export("data")
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

	err = svc.Import("data")
	
	if err != nil {
		t.Errorf("method ExportToFile returned not nil error, err => %v", err)
	}

}
