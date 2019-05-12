package response

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

type ContractedServices struct {
	CreditCards []CreditCardResponse `json:"creditCards"`
	Loans       []LoanResponse       `json:"loans"`
	Accounts    []AccountResponse    `json:"accounts"`
}

type CreditCardResponse struct {
	Name           string         `json:"categories"`
	CardNumber     string         `json:"cardNumber"`
	GoodThru       time.Time      `json:"goodThru"`
	Cvv            string         `json:"cvv"`
	CardLimit      pgtype.Numeric `json:"cardLimit"`
	InterestRate   pgtype.Numeric `json:"interestRate"`
	InterestAmount pgtype.Numeric `json:"interestAmount"`
	MonthlyCut     int            `json:"monthlyCut"`
}

type LoanResponse struct {
	Name         string         `json:"name"`
	InterestRate pgtype.Numeric `json:"interestRate"`
	Amount       pgtype.Numeric `json:"amount"`
	DuePayment   time.Time      `json:"duePayment"`
}

type AccountResponse struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Nickname string `json:"nickname"`
}
