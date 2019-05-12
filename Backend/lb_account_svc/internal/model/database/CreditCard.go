package database

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

func (CreditCard) TableName() string {
	return "life_bank_v1.credit_card"
}

type CreditCard struct {
	ID             uint           `gorm:"column:id" json:"id"`
	CardNumber     string         `gorm:"column:card_number; size:16" json:"cardNumber"`
	GoodThru       time.Time      `gorm:"column:good_thru" json:"goodThru"`
	Cvv            string         `gorm:"column:cvv; size:3" json:"cvv"`
	CardLimit      pgtype.Numeric `gorm:"column:card_limit; type:numeric" json:"cardLimit"`
	InterestRate   pgtype.Numeric `gorm:"column:interest_rate; type:numeric" json:"interestRate"`
	InterestAmount pgtype.Numeric `gorm:"column:interest_amount; type:numeric" json:"interestAmount"`
	MonthlyCut     int            `gorm:"column:monthly_cut" json:"monthlyCut"`
	CreatedDate    time.Time      `gorm:"column:created_date" json:"createdDate"`
	CreatedBy      string         `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate   time.Time      `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy     string         `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
