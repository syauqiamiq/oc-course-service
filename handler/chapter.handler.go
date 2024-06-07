package handler

import (
	"net/http"
	"ocCourseService/dto"
	"ocCourseService/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateChapterHandler(c *gin.Context) {
	var input dto.ChapterInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdChapter, err := h.service.CreateChapter(input)
	if err != nil {
		if err.Error() == "course not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatChapter(createdChapter)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) UpdateChapterHandler(c *gin.Context) {
	var input dto.UpdateChapterInput
	err := c.ShouldBindJSON(&input)

	id := c.Param("id")

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedChapter, err := h.service.UpdateChapter(id, input)
	if err != nil {
		if err.Error() == "course not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		if err.Error() == "chapter not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatChapter(updatedChapter)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetChapterHandler(c *gin.Context) {

	data, err := h.service.GetAllChapter()
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatListChapter(data)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetChapterByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetChapterByID(id)
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

	formattedResponse := dto.FormatChapter(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) DeleteChapterByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.DeleteChapterByID(id)
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

	formattedResponse := dto.FormatChapter(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
