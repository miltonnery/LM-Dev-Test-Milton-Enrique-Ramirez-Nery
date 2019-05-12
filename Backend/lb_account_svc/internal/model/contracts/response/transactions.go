package response

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

type TransactionCategoriesResponseInterface interface {
	makeResponse()
}

type TransactionCategoriesResponse struct {
	CreditCard CreditCardTransactionResponse `json:"creditCards"`
	Loan       LoanTransactionResponse       `json:"loans"`
	Account    AccountTransactionResponse    `json:"accounts"`
}

type CreditCardTransactionResponse struct {
	Name           string                `json:"categories"`
	CardNumber     string                `json:"cardNumber"`
	GoodThru       time.Time             `json:"goodThru"`
	Cvv            string                `json:"cvv"`
	CardLimit      pgtype.Numeric        `json:"cardLimit"`
	InterestRate   pgtype.Numeric        `json:"interestRate"`
	InterestAmount pgtype.Numeric        `json:"interestAmount"`
	MonthlyCut     int                   `json:"monthlyCut"`
	Transactions   []TransactionResponse `json:"transactions"`
}

type LoanTransactionResponse struct {
	Name         string                `json:"name"`
	InterestRate pgtype.Numeric        `json:"interestRate"`
	Amount       pgtype.Numeric        `json:"amount"`
	DuePayment   time.Time             `json:"duePayment"`
	Transactions []TransactionResponse `json:"transactions"`
}

type AccountTransactionResponse struct {
	Name         string                `json:"name"`
	Number       string                `json:"number"`
	Nickname     string                `json:"nickname"`
	Transactions []TransactionResponse `json:"transactions"`
}

type TransactionResponse struct {
	ID          string         `json:"id"`
	Date        time.Time      `json:"date"`
	Description string         `json:"description"`
	Amount      pgtype.Numeric `json:"amount"`
}

func (CreditCardTransactionResponse) makeResponse() {}

func (LoanTransactionResponse) makeResponse() {}

func (AccountTransactionResponse) makeResponse() {}
