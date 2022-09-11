package models

type OpsiC struct {
	Omzet        uint64 `json:"omzet" gorm:"omzet"`
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
	CreatedAt    string `json:"date" gorm:"column:date"`
}

func (OpsiC) TableName() string {
	return "transactions"
}
