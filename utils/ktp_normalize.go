package utils

import (
	"fmt"
	"micro-ocr-service/pkg/entities"
	"strings"
)

func KtpNormalize(value string) (ktp *entities.KTP, err error) {

	var (
		province         string
		city             string
		nik              string
		name             string
		placeDateOfBirth string
		gender           string
		address1         string
		address2         string
		address3         string
		address4         string
		religion         string
		marriedStatus    string
		job              string
		nationality      string
		validUntil       string
		placeOfBirth     string
		dateOfBirth      string
	)

	contents := strings.Split(strings.Replace(value, "\n\n", "\n", -1), "\n")

	for i, data := range contents {

		obj := strings.Split(data, ":")

		if len(obj) > 1 {
			data = strings.TrimSpace(obj[1])
		}

		switch i {
		case 0:
			province = data
		case 1:
			city = data
		case 2:
			nik = data
		case 3:
			name = data
		case 4:
			placeDateOfBirth = data
		case 5:
			gender = data
		case 6:
			address1 = data
		case 7:
			address2 = data
		case 8:
			address3 = data
		case 9:
			address4 = data
		case 10:
			religion = data
		case 11:
			marriedStatus = data
		case 12:
			job = data
		case 13:
			nationality = data
		case 14:
			validUntil = data
		}
	}

	fullAddress := fullAddressCast(address1, address2, address3, address4)
	placeOfBirth, dateOfBirth = placeDateOfBirthCast(placeDateOfBirth)
	gender = genderCast(gender)

	ktp = &entities.KTP{
		Province:      province,
		City:          city,
		NIK:           nik,
		Name:          name,
		PlaceOfBirth:  placeOfBirth,
		DateOfBirth:   dateOfBirth,
		Gender:        gender,
		FullAddress:   fullAddress,
		Religion:      religion,
		MarriedStatus: marriedStatus,
		Job:           job,
		Nationality:   nationality,
		ValidUntil:    validUntil,
	}

	return ktp, err
}

func genderCast(text string) string {
	femaleGender := strings.Contains(strings.ToLower(text), "perempuan")
	maleGender := strings.Contains(strings.ToLower(text), "laki-laki")
	if femaleGender {
		return "PEREMPUAN"
	}
	if maleGender {
		return "LAKI-LAKI"
	}

	return text
}

func placeDateOfBirthCast(text string) (string, string) {
	pd := strings.Split(text, ",")
	if len(pd) > 1 {
		return strings.TrimSpace(pd[0]), strings.TrimSpace(pd[len(pd)-1])
	}

	return text, ""
}

func fullAddressCast(address1, address2, address3, address4 string) string {
	return fmt.Sprintf("%s, %s, %s, %s", address1, address2, address3, address4)
}
