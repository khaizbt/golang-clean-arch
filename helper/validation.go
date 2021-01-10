package helper

import (
	"encoding/base64"
	"os"

	"github.com/go-playground/validator/v10"
)

// "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //Type data fleksibel
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {

	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func UploadImage(fileName string, content string) (bool, error) {
	decode, err := base64.StdEncoding.DecodeString(content)

	if err != nil {
		return false, err
	}
	file, err := os.Create(fileName)
	defer file.Close()
	_, err = file.Write(decode)
	if err != nil {
		return false, err
	}

	_, err = os.Stat(fileName) //Cek apakah file ada?

	if os.IsExist(err) {
		err := os.Remove(fileName)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
