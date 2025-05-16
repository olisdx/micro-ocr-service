package entities

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model

	ID         uint64           `gorm:"primaryKey;AUTO_INCREMENT"`
	Uuid       string           `gorm:"type:varchar(64); index;"`
	EntityName string           `gorm:"type:varchar(255);not null; index;"`
	BrandName  string           `gorm:"type:varchar(255);not null; index;"`
	Type       TypeMerchant     `gorm:"type:enum('personal','corporate');default:'personal'; not null"`
	Criteria   CriteriaMerchant `gorm:"type:enum('micro','small','medium','enterprise');default:'micro'; not null"`

	Email   string `gorm:"type:varchar(100); index; " `
	Address string `gorm:"type:varchar(255)"`
	Phone   string `gorm:"type:varchar(100); index;"`

	// legality
	EstablishmentDate string `gorm:"type:varchar(50)"`
	OwnerName         string `gorm:"type:varchar(255)"`
	KtpNumber         string `gorm:"type:varchar(50);"`
	NpwpNumber        string `gorm:"type:varchar(50);"`
	NibNumber         string `gorm:"type:varchar(50);"` // required if type is corporate

	// file
	KtpFile                string `gorm:"type:varchar(100);"`
	NpwpFile               string `gorm:"type:varchar(100);"`
	NibFile                string `gorm:"type:varchar(100);"` // required if type is corporate
	OperationalOfficePhoto string `gorm:"type:varchar(100);"`
	StatementLetter        string `gorm:"type:varchar(100);"` // surat pernyataan
	BankAccountBook        string `gorm:"type:varchar(100);"` // buku rekening
	RegistrationForm       string `gorm:"type:varchar(100);"`
	Agreement              string `gorm:"type:varchar(100);"` // Perjanjian Kerjasama
	Logo                   string `gorm:"type:varchar(100);"`
	StaticQrisFile         string `gorm:"type:varchar(100);"`

	// config
	BankName      string `gorm:"type:varchar(100);"`
	AccountName   string `gorm:"type:varchar(255);"`
	AccountNumber string `gorm:"type:varchar(100);"`
	EnableQris    bool   `gorm:"type:bool; default:true"`
	Mdr           string `gorm:"type:varchar(5);"` // merchant discount rate

	IsFavorite bool `gorm:"type:bool; default:false"`

	IsInactive bool `gorm:"type:bool; default:false"`
	IsBlocked  bool `gorm:"type:bool; default:false"`

	/*
			merchant personal account number;
		    a. 8 digit pertama diisi dengan NNS (National Numbering System) milik Singapay.
		    b. 1 digit diisi dengan kode Source of Fund, yang terdiri dari:
		        i. 0, jika tidak dapat dispesifikasikan.
		        ii. 1, untuk sumber dana dari akun debit.
		        iii. 2, untuk sumber dana dari akun kredit.
		        iv. 3, untuk sumber dana dari uang elektronik.
		    c. 9 digit berikutnya diisi dengan sembilan digit terakhir Merchant ID.
		    d. 1 digit terakhir diisi dengan Check-Digit Luhn.
	*/
	Mpan string `gorm:"type:varchar(255);"`

	VerifiedAt      sql.NullTime // this is not null if verification approved
	EmailVerifiedAt sql.NullTime

	MerchantCategoryCodeID uint8
	UserID                 uint64

	// Foo uint8 `json:"merchant_category_code_id"`

	// MerchantCategoryCode []MerchantCategoryCode
	// MerchantCategoryCode MerchantCategoryCode `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// MerchantCategoryCode MerchantCategoryCode `gorm:"ForeignKey:MerchantCategoryCodeID;"`

	// User                 User                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type TypeMerchant string

const (
	PersonalTypeMerchant  TypeMerchant = "personal"
	CorporateTypeMerchant TypeMerchant = "corporate"
)

func (st *TypeMerchant) Scan(value interface{}) error {
	*st = TypeMerchant(value.([]byte))
	return nil
}

func (st TypeMerchant) Value() (driver.Value, error) {
	return string(st), nil
}

type CriteriaMerchant string

const (
	MicroCriteriaMerchant      CriteriaMerchant = "micro"
	SmallCriteriaMerchant      CriteriaMerchant = "small"
	MediumCriteriaMerchant     CriteriaMerchant = "medium"
	EnterpriseCriteriaMerchant CriteriaMerchant = "enterprise"
)

func (st *CriteriaMerchant) Scan(value interface{}) error {
	*st = CriteriaMerchant(value.([]byte))
	return nil
}

func (st CriteriaMerchant) Value() (driver.Value, error) {
	return string(st), nil
}
