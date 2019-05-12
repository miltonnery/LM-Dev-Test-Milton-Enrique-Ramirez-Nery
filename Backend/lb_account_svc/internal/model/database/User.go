package database

import "time"

func (User) TableName() string {
	return "life_bank_v1.user"
}

type User struct {
	ID           uint 	   `gorm:"column:id" json:"id"`
	RoleID       uint 	   `gorm:"column:role" json:"roleId"`
	Role         Role
	Username     string    `gorm:"column:username; size:20" json:"username"`
	Password     string    `gorm:"column:password; size:100" json:"password"`
	Active       bool      `gorm:"column:active" json:"Active"`
	CreatedDate  time.Time `gorm:"column:created_date" json:"createdDate"`
	CreatedBy    string    `gorm:"column:created_by; size:50" json:"createdBy"`
	ModifiedDate time.Time `gorm:"column:modified_date" json:"modifiedDate"`
	ModifiedBy   string    `gorm:"column:modified_by; size:50" json:"modifiedBy"`
}
