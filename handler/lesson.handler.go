package handler

import (
	"net/http"
	"ocCourseService/dto"
	"ocCourseService/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateLessonHandler(c *gin.Context) {
	var input dto.LessonInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdLesson, err := h.service.CreateLesson(input)
	if err != nil {
		if err.Error() == "chapter not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatLesson(createdLesson)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) UpdateLessonHandler(c *gin.Context) {
	var input dto.UpdateLessonInput
	err := c.ShouldBindJSON(&input)

	id := c.Param("id")

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedLesson, err := h.service.UpdateLesson(id, input)
	if err != nil {
		if err.Error() == "chapter not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		if err.Error() == "lesson not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatLesson(updatedLesson)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetLessonHandler(c *gin.Context) {

	chapterId := c.Query("chapter_id")
	data, err := h.service.GetAllLesson(chapterId)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatListLesson(data)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetLessonByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetLessonByID(id)
	if err != nil {
		if err.Error() == "lesson not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatLesson(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) DeleteLessonByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.DeleteLessonByID(id)
	if err != nil {
		if err.Error() == "lesson not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatLesson(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
