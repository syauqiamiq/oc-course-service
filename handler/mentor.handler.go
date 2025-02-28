package handler

import (
	"net/http"
	"ocCourseService/dto"
	"ocCourseService/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateMentorHandler(c *gin.Context) {
	var input dto.MentorInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdMentor, err := h.service.CreateMentor(input)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatMentor(createdMentor)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) UpdateMentorHandler(c *gin.Context) {
	var input dto.UpdateMentorInput
	err := c.ShouldBindJSON(&input)

	id := c.Param("id")

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedMentor, err := h.service.UpdateMentor(id, input)
	if err != nil {
		if err.Error() == "mentor not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatMentor(updatedMentor)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetMentorHandler(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	search := c.Query("search")
	data, metaData, err := h.service.GetAllMentor(pageSize, page, search)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatListMentor(data)
	response := helper.APIResponseWithPagination("success", http.StatusOK, "Success", metaData, formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetMentorByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetMentorByID(id)
	if err != nil {
		if err.Error() == "mentor not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatMentor(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) DeleteMentorByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.DeleteMentorByID(id)
	if err != nil {
		if err.Error() == "mentor not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatMentor(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
