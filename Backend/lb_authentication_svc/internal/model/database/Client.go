package database

import "time"

func (Client) TableName() string {
	return "life_bank_v1.client"
}

type Client struct {
	ID           uint 	   `gorm:"column:id" json:"id"`
	UserInfoID   uint 	   `gorm:"column:user_info" json:"userInfo"`
	UserInfo     User
	FirstName    string    `gorm:"column:first_name; size:50" json:"firstName"`
	LastName     string    `gorm:"column:last_name; size:50" json:"lastName"`
	NationalId   string    `gorm:"column:national_id; size:20" json:"nationalId"`
	Email        string    `gorm:"column:email; size:30" json:"email"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
	Contracts []Contract
}
