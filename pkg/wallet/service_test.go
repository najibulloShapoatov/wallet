package wallet

import (

	"testing"
)


func TestService_FindAccountByID_success_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(1)

	if err != nil{
		t.Errorf("method returned not nil error, account => %v", account)
	}

}
func TestService_FindAccountByID_notFound_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+992000000001")

	account, err := svc.FindAccountByID(2)

	if err == nil{
		t.Errorf("method returned nil error, account => %v", account)
	}

}

func TestService_Reject_success_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+992000000001")
	account, err := svc.FindAccountByID(1)

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	account.Balance +=100_00
	payment, err := svc.Pay(account.ID, 10_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	err = svc.Reject(pay.ID)

	if err != nil{
		t.Errorf("method Reject returned not nil error, pay => %v", pay)
	}



}

func TestService_Reject_fail_user(t *testing.T){
	var svc Service
	svc.RegisterAccount("+992000000001")
	account, err := svc.FindAccountByID(1)

	if err != nil{
		t.Errorf("method RegisterAccount returned not nil error, account => %v", account)
	}

	account.Balance +=100_00
	payment, err := svc.Pay(account.ID, 10_00,"Cafe")

	if err != nil{
		t.Errorf("method Pay returned not nil error, account => %v", account)
	}

	pay, err := svc.FindPaymentByID(payment.ID)

	if err != nil{
		t.Errorf("method FindPaymentByID returned not nil error, payment => %v", payment)
	}

	err = svc.Reject(pay.ID+"uu")

	if err == nil{
		t.Errorf("method Reject returned not nil error, pay => %v", pay)
	}



}