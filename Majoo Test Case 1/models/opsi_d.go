package models

type OpsiD struct {
	Omzet        uint64 `json:"omzet" gorm:"omzet"`
	MerchantName string `json:"merchant_name" gorm:"column:merchant_name"`
	OutletName   string `json:"outlet_name" gorm:"column:outlet_name"`
	CreatedAt    string `json:"date" gorm:"column:date"`
}

func (OpsiD) TableName() string {
	return "transactions"
}
