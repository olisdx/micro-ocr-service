package entities

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID            uint64   `gorm:"primaryKey;AUTO_INCREMENT"`
	Uuid          string   `gorm:"type:varchar(64); index;"`
	Name          string   `gorm:"type:varchar(100);not null; index;"`
	Email         string   `gorm:"type:varchar(100); index; " `
	Password      string   `gorm:"type:varchar(255)"`
	Photo         string   `gorm:"type:varchar(100)"`
	Phone         string   `gorm:"type:varchar(100); index;"`
	Pin           string   `gorm:"type:varchar(100)"`
	BalanceWallet string   `gorm:"type:varchar(100);default:0;not null"`
	IsDeactive    bool     `gorm:"type:bool;default:false"`
	IsBlocked     bool     `gorm:"type:bool;default:false"`
	Kind          KindUser `gorm:"type:enum('basic','premium');default:'basic'; not null"`
	Role          RoleUser `gorm:"type:enum('user','merchant','admin');default:'user'; not null"`
	Mid           string   `gorm:"type:varchar(100);"`
	OtpNumber     string   `gorm:"type:varchar(10);"`
	DeviceToken   string   `gorm:"type:varchar(255);"`
	RememberToken string   `gorm:"type:varchar(100);"`

	Merchants []Merchant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	OtpExpiredAt       sql.NullTime
	LinkVerifiedAt     sql.NullTime
	PhoneVerifiedAt    sql.NullTime
	EmailVerifiedAt    sql.NullTime
	IdentityVerifiedAt sql.NullTime
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type KindUser string

const (
	BasicKind   KindUser = "basic"
	PremiumKind KindUser = "premium"
)

func (st *KindUser) Scan(value interface{}) error {
	*st = KindUser(value.([]byte))
	return nil
}

func (st KindUser) Value() (driver.Value, error) {
	return string(st), nil
}

type RoleUser string

const (
	UserRole     RoleUser = "user"
	MerchantRole RoleUser = "merchant"
	AdminRole    RoleUser = "admin"
)

func (st *RoleUser) Scan(value interface{}) error {
	*st = RoleUser(value.([]byte))
	return nil
}

func (st RoleUser) Value() (driver.Value, error) {
	return string(st), nil
}
