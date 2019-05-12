package database

import "time"

func (Contract) TableName() string {
	return "life_bank_v1.contract"
}

type Contract struct {
	ID           uint      `gorm:"column:id" json:"id"`
	ClientID     uint      `gorm:"column:client" json:"clientID"`
	Client       Client
	AccountID    uint      `gorm:"column:account" json:"account"`
	Account      Account
	CreditCardID uint      `gorm:"column:credit_card" json:"creditCard"`
	CreditCard   CreditCard
	LoanID       uint      `gorm:"column:loan" json:"loan"`
	Loan         Loan
	ProductID    uint      `gorm:"column:product" json:"product"`
	Product      Product
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
	Transactions []Transaction
}
