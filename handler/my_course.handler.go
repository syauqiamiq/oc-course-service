package handler

import (
	"net/http"
	"ocCourseService/dto"
	"ocCourseService/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateMyCourseHandler(c *gin.Context) {
	var input dto.MyCourseInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdMyCourse, err := h.service.CreateMyCourse(input)
	if err != nil {
		if err.Error() == "course not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		if err.Error() == "user not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		if err.Error() == "user already take this course" {
			response := helper.APIResponse("error", http.StatusConflict, err.Error(), nil)
			c.JSON(http.StatusConflict, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatMyCourse(createdMyCourse)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}


func (h *handler) GetMyCourseHandler(c *gin.Context) {

	userId := c.Query("user_id")
	data, err := h.service.GetAllMyCourse(userId)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatListMyCourse(data)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}