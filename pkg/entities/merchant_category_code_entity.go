package entities

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type MerchantCategoryCode struct {
	gorm.Model

	ID          uint8     `gorm:"primaryKey;AUTO_INCREMENT"`
	Code        string    `gorm:"type:varchar(100); unique;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	RiskLevel   RiskLevel `gorm:"type:enum('low','medium','high');default:'low'; not null"`
	Description string    `gorm:"type:varchar(255);"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// MerchantCategoryCode []*MerchantCategoryCode `gorm:"foreignkey:MerchantCategoryCodeID"`
	// Merchant []Merchant `gorm:"foreignkey:MerchantCategoryCodeID"`
	// Merchant []Merchant `gorm:"foreignkey:MerchantCategoryCodeID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Merchant Merchant `gorm:"foreignkey:merchant_category_code_id"`
	Merchants []Merchant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Merchants  []Merchant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RiskLevel string

const (
	LowRiskLevel    RiskLevel = "low"
	MediumRiskLevel RiskLevel = "medium"
	HighRiskLevel   RiskLevel = "high"
)

func (st *RiskLevel) Scan(value interface{}) error {
	*st = RiskLevel(value.([]byte))
	return nil
}

func (st RiskLevel) Value() (driver.Value, error) {
	return string(st), nil
}
