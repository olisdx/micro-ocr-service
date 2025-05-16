package seeders

import (
	"micro-ocr-service/pkg/entities"

	"gorm.io/gorm"
)

func MerchantCategoryCodeSeeder(db *gorm.DB) {
	data := []*entities.MerchantCategoryCode{
		{
			Code:        "5411",
			Name:        "Supermarket dan toko-toko bahan makanan",
			Description: "Supermarkets/grocery stores - Supermarket dan toko-toko bahan makanan",
			RiskLevel:   "low",
		},
		{
			Code:        "5812",
			Name:        "Restoran dan tempat makan",
			Description: "Restaurants - Restoran dan tempat makan.",
			RiskLevel:   "low",
		},
		{
			Code:        "5912",
			Name:        "Apotek dan toko obat",
			RiskLevel:   "low",
			Description: "Drug stores and pharmacies - Apotek dan toko obat",
		},
		{
			Code:        "5999",
			Name:        "Toko ritel khusus dan berbagai jenis toko lainnya.",
			RiskLevel:   "low",
			Description: "Miscellaneous and specialty retail stores - Toko ritel khusus dan berbagai jenis toko lainnya.",
		},
		{
			Code:        "7299",
			Name:        "Layanan pribadi berbagai jenis",
			RiskLevel:   "low",
			Description: "Miscellaneous personal services - Layanan pribadi berbagai jenis.",
		},
		{
			Code:        "8398",
			Name:        "Organisasi amal dan layanan sosial",
			RiskLevel:   "low",
			Description: "Charitable and social service organizations - Organisasi amal dan layanan sosial.",
		},
		{
			Code:        "8999",
			Name:        "Layanan profesional",
			RiskLevel:   "low",
			Description: "Professional services - Layanan profesional (bisa mencakup berbagai jenis layanan).",
		},
		{
			Code:        "6011",
			Name:        "Lembaga keuangan",
			RiskLevel:   "low",
			Description: "Financial Institutions - Lembaga keuangan.",
		},
		{
			Code:        "7011",
			Name:        "Hotel dan motel",
			RiskLevel:   "low",
			Description: "Hotels and motels - Hotel dan motel.",
		},
		{
			Code:        "7832",
			Name:        "Bioskop",
			RiskLevel:   "low",
			Description: "Motion Picture Theaters - Bioskop.",
		},
		{
			Code:        "4111",
			Name:        "Layanan transportasi",
			RiskLevel:   "low",
			Description: "Transportation services - Layanan transportasi.",
		},
		{
			Code:        "8211",
			Name:        "Sekolah dasar dan menengah",
			RiskLevel:   "low",
			Description: "Elementary and secondary schools - Sekolah dasar dan menengah.",
		},
		{
			Code:        "8299",
			Name:        "Sekolah dan layanan pendidikan lainnya",
			RiskLevel:   "low",
			Description: "Schools and educational services - Sekolah dan layanan pendidikan lainnya.",
		},
		{
			Code:        "7221",
			Name:        "Studio Foto",
			RiskLevel:   "low",
			Description: "Photographic studios - Studio Foto",
		},
		{
			Code:        "5541",
			Name:        "Pompa bensin/stasiun pengisian bahan bakar",
			RiskLevel:   "low",
			Description: "Service stations - Pompa bensin/stasiun pengisian bahan bakar",
		},
		{
			Code:        "4722",
			Name:        "Agen perjalanan dan operator tur",
			RiskLevel:   "low",
			Description: "Travel agencies and tour operators - Agen perjalanan dan operator tur.",
		},
		{
			Code:        "7230",
			Name:        "Salon kecantikan dan pangkas rambut",
			RiskLevel:   "low",
			Description: "Beauty and barber shops - Salon kecantikan dan pangkas rambut.",
		},
		{
			Code:        "7231",
			Name:        "Tukang rambut dan gaya rambut",
			RiskLevel:   "low",
			Description: "Hairdressers, hairstylists - Tukang rambut dan gaya rambut.",
		},
		{
			Code:        "5499",
			Name:        "Toko makanan berbagai jenis",
			RiskLevel:   "low",
			Description: "Miscellaneous food stores - Toko makanan berbagai jenis.",
		},
		{
			Code:        "6381",
			Name:        "Asuransi",
			RiskLevel:   "high",
			Description: "Insurance - Asuransi",
		},
		{
			Code:        "9991",
			Name:        "Penanaman Modal",
			RiskLevel:   "high",
			Description: "Investment Capital - Penanaman Modal",
		},
		{
			Code:        "9992",
			Name:        "Cryptocurrency",
			RiskLevel:   "high",
			Description: "Cryptocurrency",
		},
	}
	db.Create(data)

	// // work for single source
	// foo := entities.MerchantCategoryCode{
	// 	Code:        "5411",
	// 	Name:        "Supermarket dan toko-toko bahan makanan",
	// 	Description: "Supermarkets/grocery stores - Supermarket dan toko-toko bahan makanan",
	// 	RiskLevel:   "low",
	// }
	// db.Create(&foo)

}

/*

 */
