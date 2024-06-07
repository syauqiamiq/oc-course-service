package handler

import (
	"net/http"
	"ocCourseService/dto"
	"ocCourseService/helper"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCourseHandler(c *gin.Context) {
	var input dto.CourseInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdCourse, err := h.service.CreateCourse(input)
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
	formattedResponse := dto.FormatCourse(createdCourse)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) UpdateCourseHandler(c *gin.Context) {
	var input dto.UpdateCourseInput
	err := c.ShouldBindJSON(&input)

	id := c.Param("id")

	if err != nil {
		formattedError := helper.FormatValidationError(err)
		response := helper.APIResponse("error", http.StatusBadRequest, "Bad Request", formattedError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCourse, err := h.service.UpdateCourse(id, input)
	if err != nil {
		if err.Error() == "course not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		if err.Error() == "mentor not found" {
			response := helper.APIResponse("error", http.StatusNotFound, err.Error(), nil)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedResponse := dto.FormatCourse(updatedCourse)
	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetCourseHandler(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	search := c.Query("search")

	data, metaData, err := h.service.GetAllCourse(pageSize, page, search)
	if err != nil {
		response := helper.APIResponse("error", http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formattedResponse := dto.FormatListCourse(data)
	response := helper.APIResponseWithPagination("success", http.StatusOK, "Success", metaData, formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) GetCourseByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetCourseByID(id)
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

	formattedResponse := dto.FormatCourse(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}

func (h *handler) DeleteCourseByIDHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.DeleteCourseByID(id)
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

	formattedResponse := dto.FormatCourse(data)

	response := helper.APIResponse("success", http.StatusOK, "Success", formattedResponse)
	c.JSON(response.Code, response)
}
