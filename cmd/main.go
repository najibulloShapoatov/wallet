package main

import (
	"sync"
	"log"

	"github.com/najibulloShapoatov/wallet/pkg/types"
	"github.com/najibulloShapoatov/wallet/pkg/wallet"
)

func main() {

	var svc wallet.Service

	account, err := svc.RegisterAccount("+992000000001")

	if err != nil {
		log.Printf("method RegisterAccount returned not nil error, account => %v", account)
	}

	err = svc.Deposit(account.ID, 10000000_00000000000)
	if err != nil {
		log.Printf("method Deposit returned not nil error, error => %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)


	for i := 0; i < 1000; i++ {
		svc.Pay(account.ID, types.Money(i), "Cafe")
	}

	var ch <-chan types.Progress
	go func() {
		defer wg.Done()
		ch = svc.SumPaymentsWithProgress()
	}()
	go func() {
		defer wg.Done()
		ch = svc.SumPaymentsWithProgress()
	}()

	wg.Wait()

	s, ok := <-ch

	if !ok {
		log.Printf(" method SumPaymentsWithProgress ok not closed => %v", ok)
	}

	log.Println("=======>>>>>", s)
}
