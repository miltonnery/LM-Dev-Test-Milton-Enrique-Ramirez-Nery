package database

import (
	"github.com/jackc/pgx/pgtype"
	"time"
)

func (Transaction) TableName() string {
	return "life_bank_v1.transaction"
}

type Transaction struct {
	ID           uint           `gorm:"column:id" json:"id"`
	TypeID       uint           `gorm:"column:type" json:"typeID"`
	Type         TransactionType
	ContractID   uint           `gorm:"column:contract" json:"contract"`
	Contract     Contract
	Identifier   string         `gorm:"column:identifier; size:100" json:"identifier"`
	Amount       pgtype.Numeric `gorm:"column:amount; type:numeric" json:"amount"`
	Description  string         `gorm:"column:description; size:100" json:"identifier"`
	CreatedDate  time.Time      `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string         `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time      `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string         `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
