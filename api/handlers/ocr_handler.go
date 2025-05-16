package handlers

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"micro-ocr-service/pkg/service"
	"micro-ocr-service/utils"
	"mime/multipart"
	"path/filepath"
	"strings"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"
)

type OcrHandler interface {
	SendKTP(c *fiber.Ctx) error
}

type ocrHandler struct {
	service   service.OcrService
	tesseract *gosseract.Client
}

func NewOcrHandler(service service.OcrService, tesseract *gosseract.Client) ocrHandler {
	return ocrHandler{service: service, tesseract: tesseract}
}

func (h *ocrHandler) SendKTP(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(utils.Failed("form file failed"))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(utils.Failed("open file failed"))
	}
	defer src.Close()

	imagePath, modifiedImg, _ := toGrayScaleIMG(src, file.Filename)

	dst, err := os.Create(imagePath)
	if err != nil {
		return c.JSON(utils.Failed("create image failed"))
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(utils.Failed(err.Error()))
	}

	if err := jpeg.Encode(dst, modifiedImg, nil); err != nil {
		return c.JSON(utils.Failed("encode image error"))
	}

	if err = h.tesseract.SetWhitelist("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789:-/, "); err != nil {
		return c.JSON(utils.Failed("whitelist error"))
	}

	if err = h.tesseract.SetPageSegMode(gosseract.PSM_AUTO); err != nil {
		return c.JSON(utils.Failed("page segment error"))
	}

	if err = h.tesseract.SetLanguage("ind"); err != nil {
		return c.JSON(utils.Failed("set language failed"))
	}

	if err = h.tesseract.SetImage(dst.Name()); err != nil {
		return c.JSON(utils.Failed("set image failed"))
	}

	text, err := h.tesseract.Text()
	if err != nil {
		return c.JSON(utils.Failed("convert to text failed"))
	}
	os.Remove(imagePath)

	ktp, err := h.service.SendKTP(text)
	if err != nil {
		c.JSON(utils.Failed(err.Error()))
	}

	return c.JSON(utils.Success(ktp))
}

func toGrayScaleIMG(src multipart.File, imagePath string) (string, image.Image, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return "", nil, err
	}

	modifiedImg := image.NewGray(img.Bounds())
	draw.Draw(modifiedImg, modifiedImg.Bounds(), img, img.Bounds().Min, draw.Src)

	ext := filepath.Ext(imagePath)
	name := strings.TrimSuffix(filepath.Base(imagePath), ext)
	newImagePath := fmt.Sprintf("%s/%s_bw%s", filepath.Dir(imagePath), name, ext)

	return newImagePath, modifiedImg, nil
}
