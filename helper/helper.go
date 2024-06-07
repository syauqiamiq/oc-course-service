package helper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type MetaData struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PageSize    int `json:"page_size"`
	Total       int `json:"total"`
}
type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(status string, code int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type ResponseWithPagination struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Meta    MetaData    `json:"meta"`
	Data    interface{} `json:"data"`
}

func APIResponseWithPagination(status string, code int, message string, metaData MetaData, data interface{}) ResponseWithPagination {
	return ResponseWithPagination{
		Status:  status,
		Code:    code,
		Message: message,
		Meta:    metaData,
		Data:    data,
	}
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("Field %s is %s", e.Field(), e.ActualTag()))
	}
	return errors
}

func RemoveSpecialCharsAndSpaces(input string) string {
	// Define a regular expression to match special characters and spaces
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Replace matched characters with an empty string
	result := reg.ReplaceAllString(input, "")

	return result
}

func QueryPaginationTemplate(pageSize string, page string) (int, int, error) {
	convertedPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		return 0, 0, err
	}

	convertedPage, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	return convertedPageSize, convertedPage, nil
}
