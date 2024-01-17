package transaction

type Transaction struct {
	Id             string
	Amount         float64
	DocumentNumber string
	Name           string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
