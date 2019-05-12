package parsers

import (
	"lb_account_svc/internal/model/contracts/response"
	"lb_account_svc/internal/model/database"
)

func ParseTransactions(transactions []database.Transaction) *response.TransactionCategoriesResponseInterface {
	var jsonResponse response.TransactionCategoriesResponseInterface

	var ar response.AccountTransactionResponse
	var lr response.LoanTransactionResponse
	var cr response.CreditCardTransactionResponse
	var trs []response.TransactionResponse

	v := transactions[0]
	if v.Contract.AccountID != 0 {
		ar.Name = v.Contract.Product.Name
		ar.Number = v.Contract.Account.Number
		ar.Nickname = v.Contract.Account.NameID
		for _, t := range transactions {
			var tr response.TransactionResponse
			tr.ID = t.Identifier
			tr.Amount = t.Amount
			tr.Date = t.CreatedDate
			tr.Description = t.Description
			trs = append(trs, tr)
		}
		ar.Transactions = trs
		jsonResponse = ar
		return &jsonResponse
	}
	if v.Contract.LoanID != 0 {
		lr.Name = v.Contract.Product.Name
		lr.DuePayment = v.Contract.Loan.DuePayment
		lr.InterestRate = v.Contract.Loan.InterestRate
		lr.Amount = v.Contract.Loan.Amount
		for _, t := range transactions {
			var tr response.TransactionResponse
			tr.ID = t.Identifier
			tr.Amount = t.Amount
			tr.Date = t.CreatedDate
			tr.Description = t.Description
			trs = append(trs, tr)
		}
		lr.Transactions = trs
		jsonResponse = lr
		return &jsonResponse
	}
	if v.Contract.CreditCardID != 0 {
		cr.Name = v.Contract.Product.Name
		cr.InterestRate = v.Contract.CreditCard.InterestRate
		cr.InterestAmount = v.Contract.CreditCard.InterestAmount
		cr.MonthlyCut = v.Contract.CreditCard.MonthlyCut
		for _, t := range transactions {
			var tr response.TransactionResponse
			tr.ID = t.Identifier
			tr.Amount = t.Amount
			tr.Date = t.CreatedDate
			tr.Description = t.Description
			trs = append(trs, tr)
		}
		cr.Transactions = trs
		jsonResponse = cr
		return &jsonResponse
	}
	return &jsonResponse
}


