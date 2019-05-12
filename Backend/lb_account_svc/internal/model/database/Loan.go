package database

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

func (Loan) TableName() string {
	return "life_bank_v1.loan"
}

type Loan struct {
	ID           uint `gorm:"column:id" json:"id"`
	InterestRate pgtype.Numeric `gorm:"column:interest_rate; type:numeric" json:"interestRate"`
	Amount       pgtype.Numeric `gorm:"column:amount; type:numeric" json:"amount"`
	DuePayment   time.Time      `gorm:"column:due_payment" json:"duePayment"`
	CreatedDate  time.Time      `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string         `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time      `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string         `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
