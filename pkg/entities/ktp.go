package entities

type KTP struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	NIK           string `json:"nik"`
	Name          string `json:"name"`
	PlaceOfBirth  string `json:"place_of_birth"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	FullAddress   string `json:"full_address"`
	Religion      string `json:"religion"`
	MarriedStatus string `json:"married_status"`
	Job           string `json:"job"`
	Nationality   string `json:"nationality"`
	ValidUntil    string `json:"valid_until"`
}

type KtpRequest struct {
	File string `json:"file" validate:"required"`
}

type KtpResponse struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	NIK           string `json:"nik"`
	Name          string `json:"name"`
	PlaceOfBirth  string `json:"place_of_birth"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	FullAddress   string `json:"full_address"`
	Religion      string `json:"religion"`
	MarriedStatus string `json:"married_status"`
	Job           string `json:"job"`
	Nationality   string `json:"nationality"`
	ValidUntil    string `json:"valid_until"`
}

func NewKtpResponse(k *KTP) KtpResponse {
	return KtpResponse{
		Province:      k.Province,
		City:          k.City,
		NIK:           k.NIK,
		Name:          k.Name,
		PlaceOfBirth:  k.PlaceOfBirth,
		DateOfBirth:   k.DateOfBirth,
		Gender:        k.Gender,
		FullAddress:   k.FullAddress,
		Religion:      k.Religion,
		MarriedStatus: k.MarriedStatus,
		Job:           k.Job,
		Nationality:   k.Nationality,
		ValidUntil:    k.ValidUntil,
	}

}
