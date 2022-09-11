package models

type Users struct {
	ID        uint64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name      *string `json:"name"`
	UserName  *string `json:"user_name" gorm:"column:user_name"`
	Password  *string `json:"-"`
	CreatedAt string  `json:"created_at" gorm:"column:created_at"`
	CreatedBy *uint64 `json:"created_by" gorm:"column:created_by"`
	UpdatedAt string  `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy *uint64 `json:"updated_by" gorm:"column:updated_by"`
}
