package database

import "time"

func (Beneficiary) TableName() string {
	return "life_bank_v1.beneficiary"
}

type Beneficiary struct {
	ID           uint      `gorm:"column:id" json:"id"`
	OwnerID      uint      `gorm:"column:owner" json:"ownerID"`
	Owner        User
	ReceiverID   uint      `gorm:"column:receiver" json:"receiverID"`
	Receiver     User
	Identifier   string    `gorm:"column:identifier" json:"identifier"`
	Email        string    `gorm:"column:email" json:"email"`
	Active       bool      `gorm:"column:active" json:"active"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
