package seeders

import (
	"math/rand"
	"micro-ocr-service/pkg/entities"
	"strconv"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func MerchantSeeder(db *gorm.DB) {

	boolVal := []bool{true, false}
	typeVal := []entities.TypeMerchant{entities.PersonalTypeMerchant, entities.CorporateTypeMerchant}
	criteriaVal := []entities.CriteriaMerchant{entities.MicroCriteriaMerchant, entities.SmallCriteriaMerchant, entities.MediumCriteriaMerchant, entities.EnterpriseCriteriaMerchant}

	var users []entities.User
	db.Find(&users)

	var mccs []uint8
	db.Model(&entities.MerchantCategoryCode{}).Pluck("id", &mccs)

	for _, val := range users {
		result := entities.Merchant{
			Uuid:       faker.UUIDHyphenated(),
			EntityName: faker.Word(),
			BrandName:  faker.Word(),
			Type:       typeVal[rand.Intn(len(typeVal))],
			Criteria:   criteriaVal[rand.Intn(len(criteriaVal))],

			Email:   faker.Email(),
			Phone:   faker.Phonenumber(),
			Address: faker.Sentence(),

			EstablishmentDate: faker.Date(),
			OwnerName:         faker.Name(),
			KtpNumber:         strconv.Itoa(rand.Int()),
			NpwpNumber:        strconv.Itoa(rand.Int()),
			NibNumber:         strconv.Itoa(rand.Int()),

			KtpFile:                "example-file.jpg",
			NpwpFile:               "example-file.jpg",
			NibFile:                "example-file.jpg",
			OperationalOfficePhoto: "example-file.jpg",
			StatementLetter:        "example-file.jpg",
			BankAccountBook:        "example-file.jpg",
			RegistrationForm:       "example-file.jpg",
			Agreement:              "example-file.jpg",
			Logo:                   "example-file.jpg",
			StaticQrisFile:         "example-file.jpg",

			BankName:      faker.CCType(),
			AccountName:   faker.Name(),
			AccountNumber: strconv.Itoa(rand.Int()),
			EnableQris:    true,
			Mdr:           "0.25",

			IsFavorite: boolVal[rand.Intn(len(boolVal))],
			IsInactive: boolVal[rand.Intn(len(boolVal))],
			IsBlocked:  boolVal[rand.Intn(len(boolVal))],
			Mpan:       strconv.Itoa(rand.Int()),

			MerchantCategoryCodeID: mccs[rand.Intn(len(mccs))],
			UserID:                 val.ID,
		}
		db.Create(&result)
	}

}

/*

 */
