package types

//Money int64
type Money int64

//PaymentCategory string
type PaymentCategory string

//PaymentStatus string
type PaymentStatus string

//Statuses
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment model
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

//Phone string
type Phone string

//Account model
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
