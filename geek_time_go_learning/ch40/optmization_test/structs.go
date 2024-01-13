package optmization

type Request struct {
	TransactionId	string	`json:"transaction_id"`
	Payload			[]int	`json:"payload"`
}

type Response struct {
	TransactionId	string	`json:"transaction_id"`
	Expression		string	`json:"expression"`
}