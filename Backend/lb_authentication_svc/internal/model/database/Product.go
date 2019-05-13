package database

import "time"

func (Product) TableName() string {
	return "life_bank_v1.product"
}

type Product struct {
	ID           uint      `gorm:"column:id" json:"id"`
	TypeID       uint      `gorm:"column:type" json:"typeId"`
	Type         ProductType
	Name         string    `gorm:"column:name; size:50" json:"name"`
	Description  string    `gorm:"column:description" json:"active"`
	Active       bool      `gorm:"column:active; size:100" json:"description"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
