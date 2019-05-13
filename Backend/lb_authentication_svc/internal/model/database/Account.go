package database

import "time"

func (Account) TableName() string {
	return "life_bank_v1.account"
}

type Account struct {
	ID           uint 	   `gorm:"column:id" json:"id"`
	Number       string    `gorm:"column:number; size:20" json:"username"`
	NameID       string    `gorm:"column:name_id; size:20" json:"password"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
