package seeders

import (
	"math/rand"
	"micro-ocr-service/pkg/entities"
	"strconv"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {

	boolVal := []bool{true, false}
	roleVal := []entities.RoleUser{entities.UserRole, entities.MerchantRole, entities.AdminRole}
	kindVal := []entities.KindUser{entities.BasicKind, entities.PremiumKind}

	// data := []*entities.User{
	// 	&entities.User{
	// 		Uuid:          faker.UUIDHyphenated(),
	// 		Name:          faker.Name(),
	// 		Email:         faker.Email(),
	// 		Password:      faker.Password(),
	// 		Photo:         "example-photo.jpg",
	// 		Phone:         faker.Phonenumber(),
	// 		Pin:           strconv.Itoa(int(rand.Int31())),
	// 		BalanceWallet: strconv.Itoa(rand.Int()),
	// 		IsDeactive:    boolVal[rand.Intn(len(boolVal))],
	// 		IsBlocked:     boolVal[rand.Intn(len(boolVal))],
	// 		Role:          roleVal[rand.Intn(len(roleVal))],
	// 		Kind:          kindVal[rand.Intn(len(kindVal))],
	// 		Mid:           faker.CCNumber(),
	// 		OtpNumber:     "123456",
	// 		DeviceToken:   faker.Password(),
	// 		RememberToken: faker.Password(),
	// 	},
	// }
	// db.Create(data)

	// work for single source
	for i := 1; i <= 5; i++ {
		result := entities.User{
			Uuid:          faker.UUIDHyphenated(),
			Name:          faker.Name(),
			Email:         faker.Email(),
			Password:      faker.Password(),
			Photo:         "example-photo.jpg",
			Phone:         faker.Phonenumber(),
			Pin:           "123456",
			BalanceWallet: strconv.Itoa(rand.Int()),
			IsDeactive:    boolVal[rand.Intn(len(boolVal))],
			IsBlocked:     boolVal[rand.Intn(len(boolVal))],
			Role:          roleVal[rand.Intn(len(roleVal))],
			Kind:          kindVal[rand.Intn(len(kindVal))],
			Mid:           faker.CCNumber(),
			OtpNumber:     "123456",
			DeviceToken:   faker.Password(),
			RememberToken: faker.Password(),
		}
		db.Create(&result)
	}

}

/*

 */
