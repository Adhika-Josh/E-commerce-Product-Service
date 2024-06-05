package entity

import "time"

type AdminDetails struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	AdminPID  string    `gorm:"column:admin_pid;not null;"`
	Name      string    `gorm:"column:name"`
	PhoneNo   string    `gorm:"column:phone_no"`
	Dob       string    `gorm:"column:dob"`
	Email     string    `gorm:"column:email"`
	UserName  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
