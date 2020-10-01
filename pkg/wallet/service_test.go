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