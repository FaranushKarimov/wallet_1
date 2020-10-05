package types

//Money int64
type Money int64

//PaymentCategory тип - string
type PaymentCategory string

//PaymentStatus тип - string
type PaymentStatus string

//Status
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment struct
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

//Phone struct
type Phone string

//Account struct
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
