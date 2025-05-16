package service

import (
	"micro-ocr-service/pkg/entities"
	"micro-ocr-service/utils"
)

type OcrService interface {
	SendKTP(value string) (entities.KtpResponse, error)
}

type ocrService struct {
}

func NewOcrService() *ocrService {
	return &ocrService{}
}

func (s *ocrService) SendKTP(value string) (entities.KtpResponse, error) {

	ktp, err := utils.KtpNormalize(value)

	return entities.NewKtpResponse(ktp), err
}
